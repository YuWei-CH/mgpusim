package cp

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sarchlab/akita/v4/mem/mem"
	"github.com/sarchlab/akita/v4/sim"
	"github.com/sarchlab/mgpusim/v3/protocol"
)

var _ = Describe("DMAEngine", func() {
	var (
		mockCtrl          *gomock.Controller
		engine            *MockEngine
		toCP              *MockPort
		toMem             *MockPort
		localModuleFinder *mem.SingleLowModuleFinder
		dmaEngine         *DMAEngine
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		engine = NewMockEngine(mockCtrl)
		toCP = NewMockPort(mockCtrl)
		toMem = NewMockPort(mockCtrl)

		localModuleFinder = new(mem.SingleLowModuleFinder)
		dmaEngine = NewDMAEngine("DMA", engine, localModuleFinder)
		dmaEngine.ToCP = toCP
		dmaEngine.ToMem = toMem
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("should stall if dma is processing max request number", func() {
		for i := 0; i < int(dmaEngine.maxRequestCount); i++ {
			srcBuf := make([]byte, 128)
			req := protocol.NewMemCopyH2DReq(5, nil, toCP, srcBuf, uint64(20+128*i))
			rqC := NewRequestCollection(req)

			dmaEngine.processingReqs = append(dmaEngine.processingReqs, rqC)
		}

		madeProgress := dmaEngine.parseFromCP(6)

		Expect(dmaEngine.toSendToMem).To(HaveLen(0))
		Expect(madeProgress).To(BeFalse())
	})

	It("should parse MemCopyH2D from CP", func() {
		srcBuf := make([]byte, 128)
		req := protocol.NewMemCopyH2DReq(5, nil, toCP, srcBuf, 20)

		toCP.EXPECT().Retrieve(sim.VTimeInSec(6)).Return(req)

		madeProgress := dmaEngine.parseFromCP(6)

		Expect(dmaEngine.processingReqs[0].superiorRequest).To(BeIdenticalTo(req))
		Expect(dmaEngine.toSendToMem).To(HaveLen(3))
		Expect(dmaEngine.toSendToMem[0].(*mem.WriteReq).Address).
			To(Equal(uint64(20)))
		Expect(dmaEngine.toSendToMem[1].(*mem.WriteReq).Address).
			To(Equal(uint64(64)))
		Expect(dmaEngine.toSendToMem[2].(*mem.WriteReq).Address).
			To(Equal(uint64(128)))
		Expect(madeProgress).To(BeTrue())
		Expect(dmaEngine.pendingReqs).To(HaveLen(3))
	})

	It("should parse MemCopyD2H from CP", func() {
		dstBuf := make([]byte, 128)
		req := protocol.NewMemCopyD2HReq(5, nil, toCP, 20, dstBuf)

		toCP.EXPECT().Retrieve(sim.VTimeInSec(6)).Return(req)

		madeProgress := dmaEngine.parseFromCP(6)

		Expect(dmaEngine.processingReqs[0].superiorRequest).To(BeIdenticalTo(req))
		Expect(dmaEngine.toSendToMem).To(HaveLen(3))
		Expect(dmaEngine.toSendToMem[0].(*mem.ReadReq).Address).
			To(Equal(uint64(20)))
		Expect(dmaEngine.toSendToMem[1].(*mem.ReadReq).Address).
			To(Equal(uint64(64)))
		Expect(dmaEngine.toSendToMem[2].(*mem.ReadReq).Address).
			To(Equal(uint64(128)))
		Expect(madeProgress).To(BeTrue())
		Expect(dmaEngine.pendingReqs).To(HaveLen(3))
	})

	It("should parse DataReady from mem", func() {
		dstBuf := make([]byte, 128)
		req := protocol.NewMemCopyD2HReq(5, nil, toCP, 20, dstBuf)
		rqC := NewRequestCollection(req)
		dmaEngine.processingReqs = append(dmaEngine.processingReqs, rqC)

		reqToBottom1 := mem.ReadReqBuilder{}.
			WithSendTime(6).
			WithSrc(toMem).
			WithAddress(20).
			WithByteSize(64).
			Build()
		reqToBottom2 := mem.ReadReqBuilder{}.
			WithSendTime(6).
			WithSrc(toMem).
			WithAddress(64).
			WithByteSize(64).
			Build()
		reqToBottom3 := mem.ReadReqBuilder{}.
			WithSendTime(6).
			WithSrc(toMem).
			WithAddress(128).
			WithByteSize(64).
			Build()
		dmaEngine.pendingReqs = append(dmaEngine.pendingReqs, reqToBottom1)
		rqC.appendSubordinateID(reqToBottom1.Meta().ID)
		dmaEngine.pendingReqs = append(dmaEngine.pendingReqs, reqToBottom2)
		rqC.appendSubordinateID(reqToBottom2.Meta().ID)
		dmaEngine.pendingReqs = append(dmaEngine.pendingReqs, reqToBottom3)
		rqC.appendSubordinateID(reqToBottom3.Meta().ID)

		dataReady := mem.DataReadyRspBuilder{}.
			WithSendTime(7).
			WithDst(toMem).
			WithRspTo(reqToBottom2.ID).
			WithData([]byte{
				1, 2, 3, 4, 5, 6, 7, 8,
				1, 2, 3, 4, 5, 6, 7, 8,
				1, 2, 3, 4, 5, 6, 7, 8,
				1, 2, 3, 4, 5, 6, 7, 8,
				1, 2, 3, 4, 5, 6, 7, 8,
				1, 2, 3, 4, 5, 6, 7, 8,
				1, 2, 3, 4, 5, 6, 7, 8,
				1, 2, 3, 4, 5, 6, 7, 8,
			}).Build()
		toMem.EXPECT().Retrieve(sim.VTimeInSec(10)).Return(dataReady)

		madeProgress := dmaEngine.parseFromMem(10)

		Expect(madeProgress).To(BeTrue())
		Expect(dmaEngine.processingReqs[0].superiorRequest).To(BeIdenticalTo(req))
		Expect(dmaEngine.processingReqs[0]).To(BeIdenticalTo(rqC))
		Expect(dmaEngine.pendingReqs).NotTo(ContainElement(reqToBottom2))
		Expect(dmaEngine.pendingReqs).To(ContainElement(reqToBottom1))
		Expect(dmaEngine.pendingReqs).To(ContainElement(reqToBottom3))
		Expect(dstBuf[44:108]).To(Equal(dataReady.Data))
	})

	It("should respond MemCopyD2H", func() {
		dstBuf := make([]byte, 128)
		req := protocol.NewMemCopyD2HReq(5, nil, toCP, 20, dstBuf)
		rqC := NewRequestCollection(req)
		dmaEngine.processingReqs = append(dmaEngine.processingReqs, rqC)

		reqToBottom2 := mem.ReadReqBuilder{}.
			WithSendTime(6).
			WithSrc(toMem).
			WithAddress(64).
			WithByteSize(64).
			Build()
		dmaEngine.pendingReqs = append(dmaEngine.pendingReqs, reqToBottom2)
		rqC.appendSubordinateID(reqToBottom2.Meta().ID)

		dataReady := mem.DataReadyRspBuilder{}.
			WithSendTime(7).
			WithDst(toMem).
			WithRspTo(reqToBottom2.ID).
			WithData([]byte{
				1, 2, 3, 4, 5, 6, 7, 8,
				1, 2, 3, 4, 5, 6, 7, 8,
				1, 2, 3, 4, 5, 6, 7, 8,
				1, 2, 3, 4, 5, 6, 7, 8,
				1, 2, 3, 4, 5, 6, 7, 8,
				1, 2, 3, 4, 5, 6, 7, 8,
				1, 2, 3, 4, 5, 6, 7, 8,
				1, 2, 3, 4, 5, 6, 7, 8,
			}).
			Build()
		toMem.EXPECT().Retrieve(sim.VTimeInSec(10)).Return(dataReady)

		madeProgress := dmaEngine.parseFromMem(10)

		Expect(madeProgress).To(BeTrue())
		Expect(dmaEngine.processingReqs).To(BeEmpty())
		Expect(dmaEngine.pendingReqs).NotTo(ContainElement(reqToBottom2))
		Expect(dstBuf[44:108]).To(Equal(dataReady.Data))
		Expect(dmaEngine.toSendToCP[0].(*sim.GeneralRsp).OriginalReq).
			To(BeIdenticalTo(req))
	})

	It("should parse Done from mem", func() {
		srcBuf := make([]byte, 128)
		req := protocol.NewMemCopyH2DReq(5, nil, toCP, srcBuf, 20)
		rqC := NewRequestCollection(req)
		dmaEngine.processingReqs = append(dmaEngine.processingReqs, rqC)

		reqToBottom1 := mem.WriteReqBuilder{}.
			WithSendTime(6).
			WithSrc(toMem).
			WithAddress(20).
			Build()
		reqToBottom2 := mem.WriteReqBuilder{}.
			WithSendTime(6).
			WithSrc(toMem).
			WithAddress(64).
			Build()
		reqToBottom3 := mem.WriteReqBuilder{}.
			WithSendTime(6).
			WithSrc(toMem).
			WithAddress(128).
			Build()

		dmaEngine.pendingReqs = append(dmaEngine.pendingReqs, reqToBottom1)
		rqC.appendSubordinateID(reqToBottom1.Meta().ID)
		dmaEngine.pendingReqs = append(dmaEngine.pendingReqs, reqToBottom2)
		rqC.appendSubordinateID(reqToBottom2.Meta().ID)
		dmaEngine.pendingReqs = append(dmaEngine.pendingReqs, reqToBottom3)
		rqC.appendSubordinateID(reqToBottom3.Meta().ID)

		done := mem.WriteDoneRspBuilder{}.
			WithSendTime(7).
			WithDst(toMem).
			WithRspTo(reqToBottom2.ID).
			Build()

		toMem.EXPECT().Retrieve(sim.VTimeInSec(10)).Return(done)

		madeProgress := dmaEngine.parseFromMem(10)

		Expect(madeProgress).To(BeTrue())
		Expect(dmaEngine.processingReqs[0].superiorRequest).To(BeIdenticalTo(req))
		Expect(dmaEngine.processingReqs[0]).To(BeIdenticalTo(rqC))
		Expect(dmaEngine.pendingReqs).NotTo(ContainElement(reqToBottom2))
		Expect(dmaEngine.pendingReqs).To(ContainElement(reqToBottom1))
		Expect(dmaEngine.pendingReqs).To(ContainElement(reqToBottom3))
	})

	It("should send MemCopyH2D to top", func() {
		srcBuf := make([]byte, 128)
		req := protocol.NewMemCopyH2DReq(5, nil, toCP, srcBuf, 20)
		rqC := NewRequestCollection(req)
		dmaEngine.processingReqs = append(dmaEngine.processingReqs, rqC)

		reqToBottom2 := mem.WriteReqBuilder{}.
			WithSendTime(6).
			WithSrc(toMem).
			WithAddress(64).
			Build()
		dmaEngine.pendingReqs = append(dmaEngine.pendingReqs, reqToBottom2)
		rqC.appendSubordinateID(reqToBottom2.Meta().ID)

		done := mem.WriteDoneRspBuilder{}.
			WithSendTime(7).
			WithDst(toMem).
			WithRspTo(reqToBottom2.ID).
			Build()

		toMem.EXPECT().Retrieve(sim.VTimeInSec(10)).Return(done)

		madeProgress := dmaEngine.parseFromMem(10)

		Expect(madeProgress).To(BeTrue())
		Expect(dmaEngine.processingReqs).To(BeEmpty())
		Expect(dmaEngine.pendingReqs).NotTo(ContainElement(reqToBottom2))
		Expect(dmaEngine.toSendToCP[0].(*sim.GeneralRsp).OriginalReq).
			To(BeIdenticalTo(req))
	})
})

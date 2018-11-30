package driver

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gitlab.com/akita/mem/vm"
	"gitlab.com/akita/mem/vm/mock_vm"

	"gitlab.com/akita/mem"
)

var _ = Describe("Driver", func() {
	var (
		mockCtrl *gomock.Controller
		driver   *Driver
		mmu      *mock_vm.MockMMU
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mmu = mock_vm.NewMockMMU(mockCtrl)

		driver = NewDriver(nil, mmu)
		driver.registerStorage(0, 4*mem.GB)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("should allocate memory", func() {
		mmu.EXPECT().CreatePage(
			&vm.Page{
				PID:      1,
				PAddr:    0,
				VAddr:    0x100000000,
				PageSize: 4096,
				Valid:    true,
			})

		ptr := driver.AllocateMemory(8)
		Expect(ptr).To(Equal(GPUPtr(0x100000000)))

		ptr = driver.AllocateMemory(24)
		Expect(ptr).To(Equal(GPUPtr(0x100000008)))
	})

	It("should allocate memory with alignment", func() {
		mmu.EXPECT().CreatePage(
			&vm.Page{
				PID:      1,
				PAddr:    0,
				VAddr:    0x100000000,
				PageSize: 4096,
				Valid:    true,
			})

		ptr := driver.AllocateMemoryWithAlignment(8, 64)
		Expect(ptr).To(Equal(GPUPtr(0x100000000)))
		Expect(driver.allocatedPages[0]).To(HaveLen(1))
		Expect(driver.memoryMasks[0]).To(HaveLen(2))

		ptr = driver.AllocateMemoryWithAlignment(8, 64)
		Expect(ptr).To(Equal(GPUPtr(0x100000040)))
		Expect(driver.allocatedPages[0]).To(HaveLen(1))
		Expect(driver.memoryMasks[0]).To(HaveLen(4))
	})

	It("should allocate memory larger than a page", func() {
		mmu.EXPECT().CreatePage(
			&vm.Page{
				PID:      1,
				PAddr:    0,
				VAddr:    0x100000000,
				PageSize: 4096,
				Valid:    true,
			})
		mmu.EXPECT().CreatePage(
			&vm.Page{
				PID:      1,
				PAddr:    0x1000,
				VAddr:    0x100001000,
				PageSize: 4096,
				Valid:    true,
			})
		mmu.EXPECT().CreatePage(
			&vm.Page{
				PID:      1,
				PAddr:    0x2000,
				VAddr:    0x100002000,
				PageSize: 4096,
				Valid:    true,
			})

		ptr := driver.AllocateMemory(8196)
		Expect(ptr).To(Equal(GPUPtr(0x100000000)))
		Expect(driver.allocatedPages[0]).To(HaveLen(3))
	})

	It("should free memory", func() {
		//ptr := driver.AllocateMemory(4)
		//ptr2 := driver.AllocateMemory(16)
		//ptr3 := driver.AllocateMemory(8)
		//ptr4 := driver.AllocateMemory(12)
		//ptr5 := driver.AllocateMemory(24)
		//
		//driver.memoryMasks[0].Chunks[5].ByteSize = 36
		//
		//driver.FreeMemory(ptr)
		//Expect(driver.memoryMasks[0].Chunks[0].Occupied).To(Equal(false))
		//Expect(driver.memoryMasks[0].Chunks[0].ByteSize).To(Equal(uint64(4)))
		//
		//driver.FreeMemory(ptr2)
		//Expect(driver.memoryMasks[0].Chunks[0].Occupied).To(Equal(false))
		//Expect(driver.memoryMasks[0].Chunks[0].ByteSize).To(Equal(uint64(20)))
		//
		//driver.FreeMemory(ptr5)
		//Expect(driver.memoryMasks[0].Chunks[3].Occupied).To(Equal(false))
		//Expect(driver.memoryMasks[0].Chunks[3].ByteSize).To(Equal(uint64(60)))
		//
		//driver.FreeMemory(ptr4)
		//Expect(driver.memoryMasks[0].Chunks[2].Occupied).To(Equal(false))
		//Expect(driver.memoryMasks[0].Chunks[2].ByteSize).To(Equal(uint64(72)))
		//
		//driver.FreeMemory(ptr3)
		//Expect(driver.memoryMasks[0].Chunks[0].Occupied).To(Equal(false))
		//Expect(len(driver.memoryMasks[0].Chunks)).To(Equal(1))
	})

})

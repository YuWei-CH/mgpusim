	.text
	.hsa_code_object_version 2,1
	.hsa_code_object_isa 8,0,3,"AMD","AMDGPU"
	.globl	Encrypt                 ; -- Begin function Encrypt
	.p2align	8
	.type	Encrypt,@function
	.amdgpu_hsa_kernel Encrypt
Encrypt:                                ; @Encrypt
	.amd_kernel_code_t
		amd_code_version_major = 1
		amd_code_version_minor = 1
		amd_machine_kind = 1
		amd_machine_version_major = 8
		amd_machine_version_minor = 0
		amd_machine_version_stepping = 3
		kernel_code_entry_byte_offset = 256
		kernel_code_prefetch_byte_size = 0
		max_scratch_backing_memory_byte_size = 0
		granulated_workitem_vgpr_count = 4
		granulated_wavefront_sgpr_count = 1
		priority = 0
		float_mode = 192
		priv = 0
		enable_dx10_clamp = 1
		debug_mode = 0
		enable_ieee_mode = 1
		enable_sgpr_private_segment_wave_byte_offset = 0
		user_sgpr_count = 8
		enable_trap_handler = 1
		enable_sgpr_workgroup_id_x = 1
		enable_sgpr_workgroup_id_y = 0
		enable_sgpr_workgroup_id_z = 0
		enable_sgpr_workgroup_info = 0
		enable_vgpr_workitem_id = 0
		enable_exception_msb = 0
		granulated_lds_size = 0
		enable_exception = 0
		enable_sgpr_private_segment_buffer = 1
		enable_sgpr_dispatch_ptr = 1
		enable_sgpr_queue_ptr = 0
		enable_sgpr_kernarg_segment_ptr = 1
		enable_sgpr_dispatch_id = 0
		enable_sgpr_flat_scratch_init = 0
		enable_sgpr_private_segment_size = 0
		enable_sgpr_grid_workgroup_count_x = 0
		enable_sgpr_grid_workgroup_count_y = 0
		enable_sgpr_grid_workgroup_count_z = 0
		enable_ordered_append_gds = 0
		private_element_size = 1
		is_ptr64 = 1
		is_dynamic_callstack = 0
		is_debug_enabled = 0
		is_xnack_enabled = 0
		workitem_private_segment_byte_size = 0
		workgroup_group_segment_byte_size = 0
		gds_segment_byte_size = 0
		kernarg_segment_byte_size = 56
		workgroup_fbarrier_count = 0
		wavefront_sgpr_count = 15
		workitem_vgpr_count = 19
		reserved_vgpr_first = 0
		reserved_vgpr_count = 0
		reserved_sgpr_first = 0
		reserved_sgpr_count = 0
		debug_wavefront_private_segment_offset_sgpr = 0
		debug_private_segment_buffer_sgpr = 0
		kernarg_segment_alignment = 4
		group_segment_alignment = 4
		private_segment_alignment = 4
		wavefront_size = 6
		call_convention = -1
		runtime_loader_kernel_symbol = 0
	.end_amd_kernel_code_t
; BB#0:
	s_load_dword s4, s[4:5], 0x4
	s_load_dwordx2 s[0:1], s[6:7], 0x0
	s_load_dwordx2 s[2:3], s[6:7], 0x8
	s_load_dword s5, s[6:7], 0x18
	s_mov_b32 s6, 0xffff
	s_waitcnt lgkmcnt(0)
	s_and_b32 s4, s4, s6
	s_mul_i32 s8, s8, s4
	v_mov_b32_e32 v1, s1
	v_add_u32_e32 v0, vcc, s5, v0
	v_add_u32_e32 v0, vcc, s8, v0
	v_lshlrev_b32_e32 v0, 4, v0
	v_add_u32_e32 v4, vcc, s0, v0
	v_ashrrev_i32_e32 v2, 31, v0
	v_addc_u32_e32 v5, vcc, v1, v2, vcc
	flat_load_dwordx4 v[0:3], v[4:5]
	s_load_dwordx4 s[0:3], s[2:3], 0x0
	s_waitcnt lgkmcnt(0)
	s_lshr_b32 s4, s0, 16
	s_lshr_b32 s5, s0, 8
	s_lshr_b32 s6, s1, 16
	s_lshr_b32 s7, s1, 8
	s_lshr_b32 s8, s2, 16
	s_lshr_b32 s9, s2, 8
	s_lshr_b32 s10, s3, 16
	s_lshr_b32 s11, s3, 8
	v_mov_b32_e32 v8, s2
	v_mov_b32_e32 v9, s3
	v_mov_b32_e32 v10, s4
	v_mov_b32_e32 v11, s5
	v_mov_b32_e32 v12, s6
	v_mov_b32_e32 v13, s7
	s_lshr_b32 s12, s0, 24
	v_mov_b32_e32 v6, s0
	s_lshr_b32 s0, s1, 24
	v_mov_b32_e32 v7, s1
	s_lshr_b32 s1, s2, 24
	s_lshr_b32 s2, s3, 24
	v_mov_b32_e32 v14, s8
	v_mov_b32_e32 v15, s9
	v_mov_b32_e32 v16, s10
	v_mov_b32_e32 v17, s11
	s_waitcnt vmcnt(0)
	v_xor_b32_e32 v18, s12, v0
	v_xor_b32_sdwa v10, v10, v0 dst_sel:BYTE_1 dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:BYTE_1
	v_xor_b32_sdwa v11, v11, v0 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:WORD_1
	v_xor_b32_sdwa v0, v6, v0 dst_sel:BYTE_1 dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:BYTE_3
	v_xor_b32_e32 v6, s0, v1
	v_xor_b32_sdwa v12, v12, v1 dst_sel:BYTE_1 dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:BYTE_1
	v_xor_b32_sdwa v13, v13, v1 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:WORD_1
	v_xor_b32_sdwa v1, v7, v1 dst_sel:BYTE_1 dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:BYTE_3
	v_xor_b32_sdwa v7, v14, v2 dst_sel:BYTE_1 dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:BYTE_1
	v_xor_b32_sdwa v14, v15, v2 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:WORD_1
	v_xor_b32_sdwa v8, v8, v2 dst_sel:BYTE_1 dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:BYTE_3
	v_xor_b32_e32 v2, s1, v2
	v_xor_b32_sdwa v15, v16, v3 dst_sel:BYTE_1 dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:BYTE_1
	v_xor_b32_sdwa v16, v17, v3 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:WORD_1
	v_xor_b32_sdwa v9, v9, v3 dst_sel:BYTE_1 dst_unused:UNUSED_PAD src0_sel:DWORD src1_sel:BYTE_3
	v_xor_b32_e32 v3, s2, v3
	v_or_b32_sdwa v2, v2, v7 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:BYTE_0 src1_sel:DWORD
	v_or_b32_sdwa v8, v14, v8 dst_sel:WORD_1 dst_unused:UNUSED_PAD src0_sel:BYTE_0 src1_sel:DWORD
	v_or_b32_sdwa v9, v16, v9 dst_sel:WORD_1 dst_unused:UNUSED_PAD src0_sel:BYTE_0 src1_sel:DWORD
	v_or_b32_sdwa v3, v3, v15 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:BYTE_0 src1_sel:DWORD
	v_or_b32_sdwa v1, v13, v1 dst_sel:WORD_1 dst_unused:UNUSED_PAD src0_sel:BYTE_0 src1_sel:DWORD
	v_or_b32_sdwa v6, v6, v12 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:BYTE_0 src1_sel:DWORD
	v_or_b32_sdwa v0, v11, v0 dst_sel:WORD_1 dst_unused:UNUSED_PAD src0_sel:BYTE_0 src1_sel:DWORD
	v_or_b32_sdwa v7, v18, v10 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:BYTE_0 src1_sel:DWORD
	v_or_b32_sdwa v3, v3, v9 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:WORD_0 src1_sel:DWORD
	v_or_b32_sdwa v2, v2, v8 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:WORD_0 src1_sel:DWORD
	v_or_b32_sdwa v1, v6, v1 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:WORD_0 src1_sel:DWORD
	v_or_b32_sdwa v0, v7, v0 dst_sel:DWORD dst_unused:UNUSED_PAD src0_sel:WORD_0 src1_sel:DWORD
	flat_store_dwordx4 v[4:5], v[0:3]
	s_endpgm
.Lfunc_end0:
	.size	Encrypt, .Lfunc_end0-Encrypt
                                        ; -- End function
	.section	.AMDGPU.csdata
; Kernel info:
; codeLenInByte = 420
; NumSgprs: 15
; NumVgprs: 19
; ScratchSize: 0
; FloatMode: 192
; IeeeMode: 1
; LDSByteSize: 0 bytes/workgroup (compile time only)
; SGPRBlocks: 1
; VGPRBlocks: 4
; NumSGPRsForWavesPerEU: 15
; NumVGPRsForWavesPerEU: 19
; ReservedVGPRFirst: 0
; ReservedVGPRCount: 0
; COMPUTE_PGM_RSRC2:USER_SGPR: 8
; COMPUTE_PGM_RSRC2:TRAP_HANDLER: 1
; COMPUTE_PGM_RSRC2:TGID_X_EN: 1
; COMPUTE_PGM_RSRC2:TGID_Y_EN: 0
; COMPUTE_PGM_RSRC2:TGID_Z_EN: 0
; COMPUTE_PGM_RSRC2:TIDIG_COMP_CNT: 0

	.ident	"clang version 4.0 "
	.section	".note.GNU-stack"
	.amd_amdgpu_isa "amdgcn-amd-amdhsa-opencl-gfx803"
	.amd_amdgpu_hsa_metadata
---
Version:         [ 1, 0 ]
Kernels:         
  - Name:            Encrypt
    SymbolName:      'Encrypt@kd'
    Language:        OpenCL C
    LanguageVersion: [ 1, 2 ]
    Args:            
      - Name:            input
        TypeName:        'uchar*'
        Size:            8
        Align:           8
        ValueKind:       GlobalBuffer
        ValueType:       U8
        AddrSpaceQual:   Global
        AccQual:         Default
      - Name:            expanded_key
        TypeName:        'uint*'
        Size:            8
        Align:           8
        ValueKind:       GlobalBuffer
        ValueType:       U32
        AddrSpaceQual:   Global
        AccQual:         Default
      - Name:            s
        TypeName:        'uchar*'
        Size:            8
        Align:           8
        ValueKind:       GlobalBuffer
        ValueType:       U8
        AddrSpaceQual:   Global
        AccQual:         Default
      - Size:            8
        Align:           8
        ValueKind:       HiddenGlobalOffsetX
        ValueType:       I64
      - Size:            8
        Align:           8
        ValueKind:       HiddenGlobalOffsetY
        ValueType:       I64
      - Size:            8
        Align:           8
        ValueKind:       HiddenGlobalOffsetZ
        ValueType:       I64
    CodeProps:       
      KernargSegmentSize: 56
      GroupSegmentFixedSize: 0
      PrivateSegmentFixedSize: 0
      KernargSegmentAlign: 8
      WavefrontSize:   64
      NumSGPRs:        15
      NumVGPRs:        19
      MaxFlatWorkGroupSize: 256
...

	.end_amd_amdgpu_hsa_metadata

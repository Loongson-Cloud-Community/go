// Code generated by "stringer -type=RelocType"; DO NOT EDIT.

package objabi

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[R_ADDR-1]
	_ = x[R_ADDRPOWER-2]
	_ = x[R_ADDRARM64-3]
	_ = x[R_ADDRMIPS-4]
	_ = x[R_ADDRLOONG64-5]
	_ = x[R_ADDROFF-6]
	_ = x[R_WEAKADDROFF-7]
	_ = x[R_SIZE-8]
	_ = x[R_CALL-9]
	_ = x[R_CALLARM-10]
	_ = x[R_CALLARM64-11]
	_ = x[R_CALLIND-12]
	_ = x[R_CALLPOWER-13]
	_ = x[R_CALLMIPS-14]
	_ = x[R_CALLLOONG64-15]
	_ = x[R_CALLRISCV-16]
	_ = x[R_CONST-17]
	_ = x[R_PCREL-18]
	_ = x[R_TLS_LE-19]
	_ = x[R_TLS_IE-20]
	_ = x[R_GOTOFF-21]
	_ = x[R_PLT0-22]
	_ = x[R_PLT1-23]
	_ = x[R_PLT2-24]
	_ = x[R_USEFIELD-25]
	_ = x[R_USETYPE-26]
	_ = x[R_METHODOFF-27]
	_ = x[R_POWER_TOC-28]
	_ = x[R_GOTPCREL-29]
	_ = x[R_JMPMIPS-30]
	_ = x[R_JMPLOONG64-31]
	_ = x[R_DWARFSECREF-32]
	_ = x[R_DWARFFILEREF-33]
	_ = x[R_ARM64_TLS_LE-34]
	_ = x[R_ARM64_TLS_IE-35]
	_ = x[R_ARM64_GOTPCREL-36]
	_ = x[R_ARM64_GOT-37]
	_ = x[R_ARM64_PCREL-38]
	_ = x[R_ARM64_LDST8-39]
	_ = x[R_ARM64_LDST32-40]
	_ = x[R_ARM64_LDST64-41]
	_ = x[R_ARM64_LDST128-42]
	_ = x[R_POWER_TLS_LE-43]
	_ = x[R_POWER_TLS_IE-44]
	_ = x[R_POWER_TLS-45]
	_ = x[R_ADDRPOWER_DS-46]
	_ = x[R_ADDRPOWER_GOT-47]
	_ = x[R_ADDRPOWER_PCREL-48]
	_ = x[R_ADDRPOWER_TOCREL-49]
	_ = x[R_ADDRPOWER_TOCREL_DS-50]
	_ = x[R_RISCV_PCREL_ITYPE-51]
	_ = x[R_RISCV_PCREL_STYPE-52]
	_ = x[R_PCRELDBL-53]
	_ = x[R_ADDRMIPSU-54]
	_ = x[R_ADDRLOONG64U-55]
	_ = x[R_ADDRMIPSTLS-56]
	_ = x[R_ADDRLOONG64TLS-57]
	_ = x[R_ADDRLOONG64TLSU-58]
	_ = x[R_ADDRCUOFF-59]
	_ = x[R_WASMIMPORT-60]
	_ = x[R_XCOFFREF-61]
}

const _RelocType_name = "R_ADDRR_ADDRPOWERR_ADDRARM64R_ADDRMIPSR_ADDRLOONG64R_ADDROFFR_WEAKADDROFFR_SIZER_CALLR_CALLARMR_CALLARM64R_CALLINDR_CALLPOWERR_CALLMIPSR_CALLLOONG64R_CALLRISCVR_CONSTR_PCRELR_TLS_LER_TLS_IER_GOTOFFR_PLT0R_PLT1R_PLT2R_USEFIELDR_USETYPER_METHODOFFR_POWER_TOCR_GOTPCRELR_JMPMIPSR_JMPLOONG64R_DWARFSECREFR_DWARFFILEREFR_ARM64_TLS_LER_ARM64_TLS_IER_ARM64_GOTPCRELR_ARM64_GOTR_ARM64_PCRELR_ARM64_LDST8R_ARM64_LDST32R_ARM64_LDST64R_ARM64_LDST128R_POWER_TLS_LER_POWER_TLS_IER_POWER_TLSR_ADDRPOWER_DSR_ADDRPOWER_GOTR_ADDRPOWER_PCRELR_ADDRPOWER_TOCRELR_ADDRPOWER_TOCREL_DSR_RISCV_PCREL_ITYPER_RISCV_PCREL_STYPER_PCRELDBLR_ADDRMIPSUR_ADDRLOONG64UR_ADDRMIPSTLSR_ADDRLOONG64TLSR_ADDRLOONG64TLSUR_ADDRCUOFFR_WASMIMPORTR_XCOFFREF"

var _RelocType_index = [...]uint16{0, 6, 17, 28, 38, 51, 60, 73, 79, 85, 94, 105, 114, 125, 135, 148, 159, 166, 173, 181, 189, 197, 203, 209, 215, 225, 234, 245, 256, 266, 275, 287, 300, 314, 328, 342, 358, 369, 382, 395, 409, 423, 438, 452, 466, 477, 491, 506, 523, 541, 562, 581, 600, 610, 621, 635, 648, 664, 681, 692, 704, 714}

func (i RelocType) String() string {
	i -= 1
	if i < 0 || i >= RelocType(len(_RelocType_index)-1) {
		return "RelocType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _RelocType_name[_RelocType_index[i]:_RelocType_index[i+1]]
}

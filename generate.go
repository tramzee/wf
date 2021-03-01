package wf

//go:generate go run generators/gen_guids.go includes/fwpmu.h zguids.go
//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output zsyscall_windows.go syscall.go

//go:generate stringer -output=zlayerflags_strings.go -type=fwpmLayerFlags -trimprefix=fwpmLayerFlags
//go:generate stringer -output=zfieldtype_strings.go -type=fwpmFieldType -trimprefix=fwpmFieldtype
//go:generate stringer -output=zsublayerflags_strings.go -type=SublayerFlags -trimprefix=SublayerFlags
//go:generate stringer -output=zfilterenumtype_strings.go -type=FilterEnumType -trimprefix=FilterEnumType
//go:generate stringer -output=zfilterenumflags_strings.go -type=FilterEnumFlags -trimprefix=FilterEnumFlags
//go:generate stringer -output=zactiontype_strings.go -type=ActionType -trimprefix=ActionType
//go:generate stringer -output=zmatchtype_strings.go -type=MatchType -trimprefix=MatchType
//go:generate stringer -output=zfilterflags_strings.go -type=FilterFlags -trimprefix=FilterFlags
//go:generate stringer -output=zproviderflags_strings.go -type=ProviderFlags -trimprefix=ProviderFlags
//go:generate stringer -output=zdatatype_strings.go -type=DataType -trimprefix=DataType

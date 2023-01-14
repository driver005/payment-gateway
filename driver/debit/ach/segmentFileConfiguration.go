package ach

// SegmentFileConfiguration contains configuration setting for sorting during Segment File Creation.
//
// It is currently not defined, but can/will be expanded later and File.SegmentFile enhanced to use the
// configuration settings
type SegmentFileConfiguration struct{}

// SegmentFileConfiguration returns a new SegmentFileConfiguration with default values for non exported fields
func NewSegmentFileConfiguration() *SegmentFileConfiguration {
	sfc := &SegmentFileConfiguration{}
	return sfc
}

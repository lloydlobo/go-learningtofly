package layertopology

// LayerTopology implements for a network to describe
// the numbers of neurons per layer.
type LayerTopology struct {
	// Numbers of neurons per layer.
	Neurons uint
	// 		^ usize similar to uint | uintptr?
}

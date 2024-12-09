package main

// Design a system to compress images using different algorithms like JPEG, PNG, or GIF
type CompressionAlgorithm interface {
	Compress(data []byte) ([]byte, error)
}

type JPEGCompression struct{}

func (j *JPEGCompression) Compress(data []byte) ([]byte, error) {
	// Please Implement your own 🥰JPEG compression algorithm
	panic("IMPL ME")
}

type PNGCompression struct{}

func (p *PNGCompression) Compress(data []byte) ([]byte, error) {
	// Please Implement your own 🥰 PNG compression algorithm
	panic("IMPL ME")
}

type GIFCompression struct{}

func (g *GIFCompression) Compress(data []byte) ([]byte, error) {
	// Please Implement your own 🥰 GIF compression algorithm
	panic("IMPL ME")
}

type ImageProcessor struct {
	compressionAlgorithm CompressionAlgorithm
}

func (i *ImageProcessor) SetCompressionAlgorithm(algorithm CompressionAlgorithm) {
	i.compressionAlgorithm = algorithm
}

func (i *ImageProcessor) Process(data []byte) ([]byte, error) {
	return i.compressionAlgorithm.Compress(data)
}

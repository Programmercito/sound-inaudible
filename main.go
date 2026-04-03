package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"time"

	"github.com/ebitengine/oto/v3"
)

const (
	sampleRate    = 44100
	channelCount  = 2
	frequency     = 20.0             // 20Hz (threshold of hearing)
	pulseDuration = 500 * time.Millisecond
	interval      = 3 * time.Minute
	volume        = 0.05             // 5% de volumen
)

func main() {
	// 1. Configuración de Oto
	op := &oto.NewContextOptions{
		SampleRate:   sampleRate,
		ChannelCount: channelCount,
		Format:       oto.FormatFloat32LE,
	}

	context, readyNotifier, err := oto.NewContext(op)
	if err != nil {
		fmt.Printf("Error al inicializar el audio: %v\n", err)
		return
	}

	// Esperar a que el contexto esté listo
	<-readyNotifier

	fmt.Printf("--- Mantenedor de Parlantes Activo ---\n")
	fmt.Printf("Intervalo: %v\n", interval)
	fmt.Printf("Frecuencia: %v Hz (Inaudible)\n", frequency)
	fmt.Printf("Presiona Ctrl+C para detener.\n\n")

	// Primer pulso inmediato
	playPulse(context)

	// Iniciar ticker
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		playPulse(context)
	}
}

func playPulse(context *oto.Context) {
	fmt.Printf("[%s] Emitiendo pulso de mantenimiento...\n", time.Now().Format("15:04:05"))

	source := generateSine(frequency, pulseDuration, volume)
	player := context.NewPlayer(source)
	player.Play()

	// Esperar a que termine de reproducir
	time.Sleep(pulseDuration + 100*time.Millisecond)
	_ = player.Close()
}

func generateSine(freq float64, duration time.Duration, vol float32) *sineSource {
	numSamples := int(float64(sampleRate) * duration.Seconds())
	samples := make([]float32, numSamples*channelCount)

	for i := 0; i < numSamples; i++ {
		// Calcular valor de la onda
		val := float32(math.Sin(2 * math.Pi * freq * float64(i) / sampleRate))
		
		// Aplicar volumen y asignar a ambos canales (stereo)
		samples[i*2] = val * vol
		samples[i*2+1] = val * vol
	}

	return &sineSource{samples: samples}
}

type sineSource struct {
	samples []float32
	pos     int
}

func (s *sineSource) Read(buf []byte) (int, error) {
	if s.pos >= len(s.samples) {
		return 0, io.EOF
	}

	// Cada Float32 son 4 bytes
	n := len(buf) / 4
	if n > len(s.samples)-s.pos {
		n = len(s.samples) - s.pos
	}

	for i := 0; i < n; i++ {
		bits := math.Float32bits(s.samples[s.pos+i])
		binary.LittleEndian.PutUint32(buf[i*4:], bits)
	}

	s.pos += n
	return n * 4, nil
}

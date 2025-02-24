package repository

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// membuat context baru dengan batas waktu 5 detik
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// menjalankan fungsi yang membutuhkan context
	result := calculate(ctx, 10, 5)
	fmt.Println(result)

	func calculate(ctx context.Context, x int, y int) int {
		// melakukan perhitungan matematika sederhana
		sum := x + y

		// menunggu selama 10 detik untuk mensimulasikan suatu proses yang panjang
		select {
		case <-time.After(10 * time.Second):
			// jika waktu lebih dari 10 detik, kembalikan nilai 0
			return 0
		case <-ctx.Done():
			// jika context dibatalkan, kembalikan nilai -1
			return -1
		default:
			// jika tidak, kembalikan hasil perhitungan
			return sum
		}
	}
}
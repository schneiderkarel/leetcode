package flipping_an_image

func flipAndInvertImage(image [][]int) [][]int {
	if !isEntryImageValid(image) {
		return [][]int{}
	}

	for i, pixels := range image {
		var newImagePart []int

		for _, pixel := range pixels {
			if pixel == 0 {
				newImagePart = append([]int{1}, newImagePart...)
			}
			if pixel == 1 {
				newImagePart = append([]int{0}, newImagePart...)
			}
		}

		image[i] = newImagePart
	}

	return image
}

func isEntryImageValid(image [][]int) bool {
	if len(image) < 1 || len(image) > 20 {
		return false
	}

	for _, pixels := range image {
		if len(pixels) != len(image) {
			return false
		}

		for _, pixel := range pixels {
			if pixel == 0 || pixel == 1 {
				continue
			}

			return false
		}
	}

	return true
}

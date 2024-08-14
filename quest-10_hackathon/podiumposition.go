package piscine

func PodiumPosition(podium [][]string) [][]string {
	for i, v := range podium {
		for j, s := range podium[i+1:] {
			if s[0] < v[0] {
				podium[i], podium[i+j+1] = podium[i+j+1], podium[i]
			}
		}
	}
	return podium
}

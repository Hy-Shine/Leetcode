package go_leetcode

type FrequencyTracker struct {
	m         map[int]int
	frequency map[int]int
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{m: make(map[int]int), frequency: make(map[int]int)}
}

func (this *FrequencyTracker) Add(number int) {
	fre := this.m[number]
	if this.frequency[fre] > 0 {
		this.frequency[fre]--
	}
	this.frequency[fre+1]++
	this.m[number]++
}

func (this *FrequencyTracker) DeleteOne(number int) {
	if this.m[number] == 0 {
		return
	}
	if this.frequency[this.m[number]] > 0 {
		this.frequency[this.m[number]]--
	}
	this.frequency[this.m[number]-1]++
	this.m[number]--
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	val := this.frequency[frequency]
	return val != 0
}

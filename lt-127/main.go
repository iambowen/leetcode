package main

import "strings"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	index := 0
	for ; index < len(wordList); index++ {
		if wordList[index] == endWord {
			break
		}
	}

	if index == len(wordList) {
		return 0
	}

	// visited := make(map[string]bool, len(wordList))
	visited := make(map[string]bool, len(wordList))
	visited[beginWord] = true
	list := constructGraph(wordList, beginWord)
	queue := []string{}
	queue = list[beginWord]
	steps := 1
	for len(queue) != 0 {
		temp := queue
		queue = []string{}
		steps++
		for i := 0; i < len(temp); i++ {
			if visited[temp[i]] {
				continue
			}
			if temp[i] == endWord {
				return steps
			}
			queue = append(queue, list[temp[i]]...)
			visited[temp[i]] = true
		}
	}

	return 0
}

func isOneBitDiff(a string, b string) bool {
	diff := 0
	a1 := strings.Split(a, "")
	b1 := strings.Split(b, "")
	for i := 0; i < len(a1); i++ {
		if a1[i] != b1[i] {
			diff++
		}
	}

	return diff == 1
}

func constructGraph(wordList []string, beginWord string) map[string][]string {
	result := map[string][]string{}
	a := []string{beginWord}
	wordList = append(a, wordList...)

	for _, word := range wordList {
		temp := []string{}
		for _, anotherWord := range wordList {
			if isOneBitDiff(word, anotherWord) {
				temp = append(temp, anotherWord)
			}
		}
		result[word] = temp
	}

	return result

}
func main() {

}




total=40 //会议结束后总的反馈人数
gifts=20 //剩余礼品的数目，包括手环，书籍等礼品

(1..30).map{|x| rand(total)}.uniq().take(gifts)  //循环30次，随机30个值，去重，然后取前gifts个礼品

//[21, 24, 17, 16, 25, 22, 3, 8, 15, 39, 11, 6, 28, 30, 23] 按照顺序叫号，从书籍一直到最后的手环，发给中奖听众即可
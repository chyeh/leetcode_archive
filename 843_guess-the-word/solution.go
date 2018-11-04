package solution

import "math/rand"

type Master struct {
	secret string
}

func (this *Master) Guess(word string) int { return distance(word, this.secret) }

func distance(a, b string) int {
	var cnt int
	for i := range a {
		if a[i] != b[i] {
			cnt++
		}
	}
	return cnt
}

func findSecretWord(wordlist []string, master *Master) {
	for i := 0; i < 15; i++ {
		candidate := wordlist[rand.Intn(len(wordlist))]
		d := 6 - master.Guess(candidate)
		if d == 0 {
			return
		}
		for i := 0; i < len(wordlist); {
			if distance(candidate, wordlist[i]) != d {
				wordlist = append(wordlist[:i], wordlist[i+1:]...)
			} else {
				i++
			}
		}
	}
}

// Test Case:
/*
"ccoyyo"
["wichbx","oahwep","tpulot","eqznzs","vvmplb","eywinm","dqefpt","kmjmxr","ihkovg","trbzyb","xqulhc","bcsbfw","rwzslk","abpjhw","mpubps","viyzbc","kodlta","ckfzjh","phuepp","rokoro","nxcwmo","awvqlr","uooeon","hhfuzz","sajxgr","oxgaix","fnugyu","lkxwru","mhtrvb","xxonmg","tqxlbr","euxtzg","tjwvad","uslult","rtjosi","hsygda","vyuica","mbnagm","uinqur","pikenp","szgupv","qpxmsw","vunxdn","jahhfn","kmbeok","biywow","yvgwho","hwzodo","loffxk","xavzqd","vwzpfe","uairjw","itufkt","kaklud","jjinfa","kqbttl","zocgux","ucwjig","meesxb","uysfyc","kdfvtw","vizxrv","rpbdjh","wynohw","lhqxvx","kaadty","dxxwut","vjtskm","yrdswc","byzjxm","jeomdc","saevda","himevi","ydltnu","wrrpoc","khuopg","ooxarg","vcvfry","thaawc","bssybb","ccoyyo","ajcwbj","arwfnl","nafmtm","xoaumd","vbejda","kaefne","swcrkh","reeyhj","vmcwaf","chxitv","qkwjna","vklpkp","xfnayl","ktgmfn","xrmzzm","fgtuki","zcffuv","srxuus","pydgmq"]
10
*/

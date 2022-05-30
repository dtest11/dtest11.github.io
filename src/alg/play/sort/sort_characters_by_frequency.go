package sort

/***
451. Sort Characters By Frequency
Medium

1960

138

Add to List

Share
Given a string, sort it in decreasing order based on the frequency of characters.

Example 1:

Input:
"tree"

Output:
"eert"

Explanation:
'e' appears twice while 'r' and 't' both appear once.
So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
*/
type Record struct {
	Key     rune
	Counter int
}

func frequencySort(s string) string {
	// a  Ascall=97, z=112
	var records []Record

	var allKeys []rune
	var exist = make(map[rune]int)
	for _, v := range s {
		count, ex := exist[v]
		if ex {
			exist[v] = count + 1
		} else {
			exist[v] = 1
			allKeys = append(allKeys, v)
		}
	}

	//sort records accord ascall code value
	for i := 0; i < len(allKeys); i++ {
		for j := i + 1; j < len(allKeys) && j-1 >= 0; j-- {
			if allKeys[j-1] < allKeys[j] {
				break // sorted
			} else {
				allKeys[j-1], allKeys[j] = allKeys[j], allKeys[j-1]
			}
		}
	}
	for _, v := range allKeys {
		t := Record{
			Key:     v,
			Counter: exist[v],
		}
		records = append(records, t)
	}
	//sort records
	for i := 0; i < len(records); i++ {
		for j := i + 1; j < len(records) && j-1 >= 0; j-- {
			if records[j-1].Counter >= records[j].Counter {
				break // sorted
			} else {
				records[j-1], records[j] = records[j], records[j-1]
			}
		}
	}
	var result string
	for _, v := range records {
		counter := v.Counter
		tmp := ""
		for i := 0; i < counter; i++ {
			tmp += string(v.Key)
		}
		result += tmp

	}
	return result
}

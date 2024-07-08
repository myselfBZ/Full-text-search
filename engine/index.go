package engine

type Index map[string][]int

func (i *Index) Add(docs []Document)  {
    for _, d := range docs{
        for _, token := range analyze(d.Text){
            ids := (*i)[token]
            if ids != nil && ids[len(ids) - 1] == d.ID {
                continue 
            }
            (*i)[token] = append((*i)[token], d.ID)
        }
    }
}


func Intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	r := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

func (i *Index) Search(text string) []int  {
    var r []int 
    for _, token := range analyze(text){
        if ids, ok := (*i)[token]; ok{
            if r == nil {
                r = ids 
            } else{
                r = Intersection(r, ids)
            }
        } else {
            return nil 
        }
    }
    return r 
}

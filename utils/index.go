package utils

type Index map[string][]int

func (i *Index) Add(docs []Document)  {
    for _, d := range docs{
        for _, token := range analyze(d.Text){
            ids := i[token]
        }
    }
}

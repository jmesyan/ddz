package ddz

import (
	"reflect"
	"testing"
)

func Test_newHandContext(t *testing.T) {
	type args struct {
		cs CardSlice
	}
	tests := []struct {
		name string
		args args
		want *handContext
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newHandContext(tt.args.cs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newHandContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handContext_copy(t *testing.T) {
	type fields struct {
		ranks    RankCount
		cards    CardSlice
		reversed CardSlice
	}
	tests := []struct {
		name   string
		fields fields
		want   *handContext
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &handContext{
				ranks:    tt.fields.ranks,
				cards:    tt.fields.cards,
				reversed: tt.fields.reversed,
			}
			if got := ctx.copy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handContext.copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handContext_update(t *testing.T) {
	type fields struct {
		ranks    RankCount
		cards    CardSlice
		reversed CardSlice
	}
	type args struct {
		cs CardSlice
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *handContext
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &handContext{
				ranks:    tt.fields.ranks,
				cards:    tt.fields.cards,
				reversed: tt.fields.reversed,
			}
			if got := ctx.update(tt.args.cs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handContext.update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handContext_searchPrimal(t *testing.T) {
	type fields struct {
		ranks    RankCount
		cards    CardSlice
		reversed CardSlice
	}
	type args struct {
		toBeat    *Hand
		primalNum int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Hand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &handContext{
				ranks:    tt.fields.ranks,
				cards:    tt.fields.cards,
				reversed: tt.fields.reversed,
			}
			if got := ctx.searchPrimal(tt.args.toBeat, tt.args.primalNum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handContext.searchPrimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handContext_searchBomb(t *testing.T) {
	type fields struct {
		ranks    RankCount
		cards    CardSlice
		reversed CardSlice
	}
	type args struct {
		toBeat *Hand
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Hand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &handContext{
				ranks:    tt.fields.ranks,
				cards:    tt.fields.cards,
				reversed: tt.fields.reversed,
			}
			if got := ctx.searchBomb(tt.args.toBeat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handContext.searchBomb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handContext_searchTrioKicker(t *testing.T) {
	type fields struct {
		ranks    RankCount
		cards    CardSlice
		reversed CardSlice
	}
	type args struct {
		toBeat    *Hand
		kickerNum int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Hand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &handContext{
				ranks:    tt.fields.ranks,
				cards:    tt.fields.cards,
				reversed: tt.fields.reversed,
			}
			if got := ctx.searchTrioKicker(tt.args.toBeat, tt.args.kickerNum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handContext.searchTrioKicker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handContext_searchFourKicker(t *testing.T) {
	type fields struct {
		ranks    RankCount
		cards    CardSlice
		reversed CardSlice
	}
	type args struct {
		toBeat    *Hand
		kickerNum int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Hand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &handContext{
				ranks:    tt.fields.ranks,
				cards:    tt.fields.cards,
				reversed: tt.fields.reversed,
			}
			if got := ctx.searchFourKicker(tt.args.toBeat, tt.args.kickerNum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handContext.searchFourKicker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handContext_searchChain(t *testing.T) {
	type fields struct {
		ranks    RankCount
		cards    CardSlice
		reversed CardSlice
	}
	type args struct {
		toBeat    *Hand
		duplicate int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Hand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &handContext{
				ranks:    tt.fields.ranks,
				cards:    tt.fields.cards,
				reversed: tt.fields.reversed,
			}
			if got := ctx.searchChain(tt.args.toBeat, tt.args.duplicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handContext.searchChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextComb(t *testing.T) {
	type args struct {
		comb []int
		k    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextComb(tt.args.comb, tt.args.k, tt.args.n); got != tt.want {
				t.Errorf("nextComb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handContext_searchTrioKickerChain(t *testing.T) {
	type fields struct {
		ranks    RankCount
		cards    CardSlice
		reversed CardSlice
	}
	type args struct {
		toBeat *Hand
		kc     int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Hand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &handContext{
				ranks:    tt.fields.ranks,
				cards:    tt.fields.cards,
				reversed: tt.fields.reversed,
			}
			if got := ctx.searchTrioKickerChain(tt.args.toBeat, tt.args.kc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handContext.searchTrioKickerChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardSlice_SearchBeat(t *testing.T) {
	type args struct {
		toBeat *Hand
	}
	tests := []struct {
		name string
		cs   CardSlice
		args args
		want *Hand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cs.SearchBeat(tt.args.toBeat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CardSlice.SearchBeat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardSlice_SearchBeatList(t *testing.T) {
	type args struct {
		toBeat *Hand
	}
	tests := []struct {
		name string
		cs   *CardSlice
		args args
		want []*Hand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cs.SearchBeatList(tt.args.toBeat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CardSlice.SearchBeatList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractConsecutive(t *testing.T) {
	type args struct {
		cs        CardSlice
		duplicate int
	}
	tests := []struct {
		name  string
		args  args
		want  CardSlice
		want1 []*Hand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := extractConsecutive(tt.args.cs, tt.args.duplicate)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractConsecutive() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("extractConsecutive() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_extractNukeBombDeuce(t *testing.T) {
	type args struct {
		cs CardSlice
		rc RankCount
	}
	tests := []struct {
		name  string
		args  args
		want  CardSlice
		want1 RankCount
		want2 []*Hand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := extractNukeBombDeuce(tt.args.cs, tt.args.rc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractNukeBombDeuce() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("extractNukeBombDeuce() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("extractNukeBombDeuce() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestStandardAnalyze(t *testing.T) {
	type args struct {
		cs CardSlice
	}
	tests := []struct {
		name string
		args args
		want []*Hand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StandardAnalyze(tt.args.cs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StandardAnalyze() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handContext_findLongestConsecutive(t *testing.T) {
	type fields struct {
		ranks    RankCount
		cards    CardSlice
		reversed CardSlice
	}
	type args struct {
		duplicate int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Hand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &handContext{
				ranks:    tt.fields.ranks,
				cards:    tt.fields.cards,
				reversed: tt.fields.reversed,
			}
			if got := ctx.findLongestConsecutive(tt.args.duplicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handContext.findLongestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handContext_traverseChains(t *testing.T) {
	type fields struct {
		ranks    RankCount
		cards    CardSlice
		reversed CardSlice
	}
	type args struct {
		last      *Hand
		duplicate *int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Hand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &handContext{
				ranks:    tt.fields.ranks,
				cards:    tt.fields.cards,
				reversed: tt.fields.reversed,
			}
			if got := ctx.traverseChains(tt.args.last, tt.args.duplicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handContext.traverseChains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handContext_extractAllChains(t *testing.T) {
	type fields struct {
		ranks    RankCount
		cards    CardSlice
		reversed CardSlice
	}
	tests := []struct {
		name   string
		fields fields
		want   []*Hand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &handContext{
				ranks:    tt.fields.ranks,
				cards:    tt.fields.cards,
				reversed: tt.fields.reversed,
			}
			if got := ctx.extractAllChains(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handContext.extractAllChains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchTreeNode_copy(t *testing.T) {
	type fields struct {
		ctx    *handContext
		hand   *Hand
		weight int
	}
	tests := []struct {
		name   string
		fields fields
		want   *searchTreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &searchTreeNode{
				ctx:    tt.fields.ctx,
				hand:   tt.fields.hand,
				weight: tt.fields.weight,
			}
			if got := n.copy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchTreeNode.copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newSearchTreeNode(t *testing.T) {
	type args struct {
		ctx *handContext
	}
	tests := []struct {
		name string
		args args
		want *searchTreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newSearchTreeNode(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSearchTreeNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newSearchTree(t *testing.T) {
	type args struct {
		n *searchTreeNode
	}
	tests := []struct {
		name string
		args args
		want *searchTree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newSearchTree(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSearchTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchTree_addChild(t *testing.T) {
	type fields struct {
		node     *searchTreeNode
		parent   *searchTree
		children []*searchTree
	}
	type args struct {
		hand *Hand
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *searchTree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := &searchTree{
				node:     tt.fields.node,
				parent:   tt.fields.parent,
				children: tt.fields.children,
			}
			if got := t.addChild(tt.args.hand); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchTree.addChild() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchTree_dumpLeaves(t *testing.T) {
	type fields struct {
		node     *searchTreeNode
		parent   *searchTree
		children []*searchTree
	}
	tests := []struct {
		name   string
		fields fields
		want   []*searchTree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := &searchTree{
				node:     tt.fields.node,
				parent:   tt.fields.parent,
				children: tt.fields.children,
			}
			if got := t.dumpLeaves(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchTree.dumpLeaves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdvancedAnalyze(t *testing.T) {
	type args struct {
		cs CardSlice
	}
	tests := []struct {
		name string
		args args
		want []*Hand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AdvancedAnalyze(tt.args.cs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AdvancedAnalyze() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStandardEvaluator_Evaluate(t *testing.T) {
	type fields struct {
		Evaluator Evaluator
	}
	type args struct {
		cs CardSlice
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := StandardEvaluator{
				Evaluator: tt.fields.Evaluator,
			}
			if got := e.Evaluate(tt.args.cs); got != tt.want {
				t.Errorf("StandardEvaluator.Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdvancedEvaluator_Evaluate(t *testing.T) {
	type fields struct {
		Evaluator Evaluator
	}
	type args struct {
		cs CardSlice
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := AdvancedEvaluator{
				Evaluator: tt.fields.Evaluator,
			}
			if got := e.Evaluate(tt.args.cs); got != tt.want {
				t.Errorf("AdvancedEvaluator.Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBestBeat(t *testing.T) {
	type args struct {
		cs        CardSlice
		toBeat    *Hand
		evaluator Evaluator
	}
	tests := []struct {
		name string
		args args
		want *Hand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BestBeat(tt.args.cs, tt.args.toBeat, tt.args.evaluator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BestBeat() = %v, want %v", got, tt.want)
			}
		})
	}
}

package blockchain

import (
	"reflect"
	"sync"
	"testing"

	"github.com/NamSeungwoo/namcoin/utils"
)

type fakeDB struct {
	fakeLoadChain func() []byte
	fakeFindBlock func() []byte
}

func (f fakeDB) FindBlock(hash string) []byte {
	return f.fakeFindBlock()
}
func (fakeDB) SaveBlock(hash string, data []byte) {

}
func (fakeDB) SaveChain(data []byte) {

}
func (f fakeDB) LoadChain() []byte {
	return f.fakeLoadChain()
}
func (fakeDB) DeleteAllBlocks() {

}

func TestRecalculateDifficulty(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		blocks := []*Block{
			{PrevHash: "x", Timestamp: 120},
			{PrevHash: "x", Timestamp: 60},
			{PrevHash: "x", Timestamp: 60},
			{PrevHash: "x", Timestamp: 60},
			{PrevHash: "", Timestamp: 0},
		}
		fakeBlock := 0
		dbStorage = fakeDB{
			fakeFindBlock: func() []byte {
				defer func() {
					fakeBlock++
				}()
				return utils.ToBytes(blocks[fakeBlock])
			},
		}
		bc := &blockchain{} //{Height: 5, CurrentDifficulty: 3}
		got := recalculateDifficulty(bc)
		t.Errorf(" recalculateDifficulty should return %d got %d", bc.CurrentDifficulty+1, got)

	})
	t.Run("2", func(t *testing.T) {
		blocks := []*Block{
			{PrevHash: "x", Timestamp: 480},
			{PrevHash: "x", Timestamp: 60},
			{PrevHash: "x", Timestamp: 60},
			{PrevHash: "x", Timestamp: 60},
			{PrevHash: "", Timestamp: 0},
		}
		fakeBlock := 0
		dbStorage = fakeDB{
			fakeFindBlock: func() []byte {
				defer func() {
					fakeBlock++
				}()
				return utils.ToBytes(blocks[fakeBlock])
			},
		}
		bc := &blockchain{Height: 5, CurrentDifficulty: 3}
		got := recalculateDifficulty(bc)
		t.Errorf(" recalculateDifficulty should return %d got %d", bc.CurrentDifficulty, got)

	})
	t.Run("3", func(t *testing.T) {
		once = *new(sync.Once)
		blocks := []*Block{
			{PrevHash: "x", Timestamp: 4800},
			{PrevHash: "x", Timestamp: 60},
			{PrevHash: "x", Timestamp: 60},
			{PrevHash: "x", Timestamp: 60},
			{PrevHash: "", Timestamp: 0},
		}
		fakeBlock := 0
		dbStorage = fakeDB{
			fakeFindBlock: func() []byte {
				defer func() {
					fakeBlock++
				}()
				return utils.ToBytes(blocks[fakeBlock])
			},
		}
		bc := &blockchain{} //{Height: 5, CurrentDifficulty: 3}
		got := recalculateDifficulty(bc)
		t.Errorf(" recalculateDifficulty should return %d got %d", bc.CurrentDifficulty-1, got)
	})
}

func TestGetDifficulty(t *testing.T) {
	blocks := []*Block{
		{PrevHash: "x", Timestamp: 720},
		{PrevHash: "x", Timestamp: 60},
		{PrevHash: "x", Timestamp: 60},
		{PrevHash: "x", Timestamp: 60},
		{PrevHash: "", Timestamp: 0},
	}
	fakeBlock := 0
	dbStorage = fakeDB{
		fakeFindBlock: func() []byte {
			defer func() {
				fakeBlock++
			}()
			return utils.ToBytes(blocks[fakeBlock])
		},
	}
	type test struct {
		height int
		want   int
	}
	tests := []test{
		{height: 0, want: defaultDifficulty},
		{height: 2, want: defaultDifficulty},
		{height: 5, want: 3},
	}
	for _, tc := range tests {
		bc := &blockchain{Height: tc.height, CurrentDifficulty: defaultDifficulty}
		got := getDifficulty(bc)
		if got != tc.want {
			t.Errorf("getDifficulty() should return %d got %d", tc.want, got)
		}
	}
}

func TestBlockchain(t *testing.T) {
	t.Run("Should create blockchain", func(t *testing.T) {
		dbStorage = fakeDB{
			fakeLoadChain: func() []byte {
				return nil
			},
		}
		bc := Blockchain()
		if bc.Height != 1 {
			t.Error("Blockchain() should create a blockchain")
		}
	})
	t.Run("Should restore blockchain", func(t *testing.T) {
		once = *new(sync.Once)
		dbStorage = fakeDB{
			fakeLoadChain: func() []byte {
				bc := &blockchain{Height: 2, NewestHash: "xxx", CurrentDifficulty: 1}
				return utils.ToBytes(bc)
			},
		}
		bc := Blockchain()
		if bc.Height != 2 {
			t.Errorf("Blockchain() should restore a blockchain with a height of %d, got %d", 2, bc.Height)
		}
	})
}

func TestBlocks(t *testing.T) { // TestBlocks 테스트 함수는 한번 호출되지만, Blocks의 FindBlock으로 인해(loop) fakeFindBlock은 여러번 호출이 된다.
	blocks := []*Block{
		{PrevHash: "x"},
		{PrevHash: ""},
	}
	fakeBlock := 0 // fakeFindBlock이 여러번 호출, 호출 횟수 상태 저장하기 위한 변수.. 전역변수 역할
	dbStorage = fakeDB{
		fakeFindBlock: func() []byte {
			defer func() {
				fakeBlock++
			}()
			return utils.ToBytes(blocks[fakeBlock])
		},
	}
	bc := &blockchain{}
	blocksResult := Blocks(bc)
	if reflect.TypeOf(blocksResult) != reflect.TypeOf([]*Block{}) {
		t.Error("Blocks() should return a slice of blocks")
	}

}

func TestFindTx(t *testing.T) {
	t.Run("Tx not found", func(t *testing.T) {
		dbStorage = fakeDB{
			fakeFindBlock: func() []byte {
				b := &Block{
					Height:       2,
					Transactions: []*Tx{},
				}

				return utils.ToBytes(b)
			},
		}
		tx := FindTx(&blockchain{}, "test")
		if tx != nil {
			t.Error("Tx should be not found.")
		}
	})
	t.Run("Tx should be found", func(t *testing.T) {
		dbStorage = fakeDB{
			fakeFindBlock: func() []byte {
				b := &Block{
					Height: 2,
					Transactions: []*Tx{
						{Id: "test"},
					},
				}

				return utils.ToBytes(b)
			},
		}
		tx := FindTx(&blockchain{}, "test")
		if tx == nil {
			t.Error("Tx should be found.")
		}
	})
}

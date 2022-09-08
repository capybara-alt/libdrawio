package libdrawio_test

import (
	"testing"

	"github.com/capybara-alt/libdrawio/libdrawio"
)

func TestDecompress(t *testing.T) {
	tests := "lJI9b8MgGIR/DWMlA1WaDl5i5WNIl2ZIOyLz1iDxJUJip7++ChCMoizdfA9nOO4F0U5PW8+c+LAcFCINn1pECMbNAhFy09ebXuJlkoOXPBpmeZC/EFGT2VlyOFWmYK0K0tWot8ZAHyrCvLdjbfmxqj7JsQEe5KFn6pEdJQ8iJiZvM92BHEQ6DS/eE9csGXPqk2DcjgXQNaKdtzakLz11oG71pNvPFkQ2ee2ZrwTzYMJ/foRud/naD+bTuOG4N99yp8lL3u3C1DleOwcP19SCt2fDgbcNoqtRyAAHx3poR88coisRtGoxoqvHSPddwQeYCsihtmA1BH9FpIlrZcjxVeDXrMZ76/hepyiNLzJhcbhD2bDuoRzzrBpENvMgkn1+sHT9BwAA//8BAAD//w=="
	mxGraphModel, err := libdrawio.Decompress(tests)
	if err != nil {
		t.Fail()
	}

	if len(mxGraphModel.Content.MxCells) < 1 {
		t.Fail()
	}

	if mxGraphModel.Content.MxCells[0].Id != "0" {
		t.Fail()
	}

	if mxGraphModel.Content.MxCells[1].Id != "1" {
		t.Fail()
	}

	if mxGraphModel.Content.MxCells[2].Id != "eCHvXLgnRnpgWLnYiHm2-1" {
		t.Fail()
	}
}

func TestCompress(t *testing.T) {
	tests := "lJI9b8MgGIR/DWMlA1WaDl5i5WNIl2ZIOyLz1iDxJUJip7++ChCMoizdfA9nOO4F0U5PW8+c+LAcFCINn1pECMbNAhFy09ebXuJlkoOXPBpmeZC/EFGT2VlyOFWmYK0K0tWot8ZAHyrCvLdjbfmxqj7JsQEe5KFn6pEdJQ8iJiZvM92BHEQ6DS/eE9csGXPqk2DcjgXQNaKdtzakLz11oG71pNvPFkQ2ee2ZrwTzYMJ/foRud/naD+bTuOG4N99yp8lL3u3C1DleOwcP19SCt2fDgbcNoqtRyAAHx3poR88coisRtGoxoqvHSPddwQeYCsihtmA1BH9FpIlrZcjxVeDXrMZ76/hepyiNLzJhcbhD2bDuoRzzrBpENvMgkn1+sHT9BwAA//8BAAD//w=="
	mxGraphModel, err := libdrawio.Decompress(tests)
	if err != nil {
		t.Fail()
	}

	compressed, err := mxGraphModel.Compress()
	if err != nil {
		t.Fail()
	}

	if compressed != tests {
		t.Fail()
	}
}

func TestGetIds(t *testing.T) {
	tests := "lJI9b8MgGIR/DWMlA1WaDl5i5WNIl2ZIOyLz1iDxJUJip7++ChCMoizdfA9nOO4F0U5PW8+c+LAcFCINn1pECMbNAhFy09ebXuJlkoOXPBpmeZC/EFGT2VlyOFWmYK0K0tWot8ZAHyrCvLdjbfmxqj7JsQEe5KFn6pEdJQ8iJiZvM92BHEQ6DS/eE9csGXPqk2DcjgXQNaKdtzakLz11oG71pNvPFkQ2ee2ZrwTzYMJ/foRud/naD+bTuOG4N99yp8lL3u3C1DleOwcP19SCt2fDgbcNoqtRyAAHx3poR88coisRtGoxoqvHSPddwQeYCsihtmA1BH9FpIlrZcjxVeDXrMZ76/hepyiNLzJhcbhD2bDuoRzzrBpENvMgkn1+sHT9BwAA//8BAAD//w=="
	mxGraphModel, err := libdrawio.Decompress(tests)
	if err != nil {
		t.Fail()
	}

	ids := mxGraphModel.GetIds()
	if ids[0] != "0" {
		t.Fail()
	}

	if ids[1] != "1" {
		t.Fail()
	}

	if ids[2] != "eCHvXLgnRnpgWLnYiHm2-1" {
		t.Fail()
	}
}

func TestGetIdsFromObjectAndMxCells(t *testing.T) {
	tests := "xZVRT4MwEMc/DY8mtExkr5vTxeiLi9l87MZZagolXSdsn94ixwDHFIzGJ3r/O9rr7+5Sx5vG+a1mafSgQpAOdcPc8a4dSglxffsplH2pBCQoBa5FiEG1sBAHQNFFdSdC2LYCjVLSiLQtblSSwMa0NKa1ytphL0q2T00ZhxNhsWHyVF2K0ER4C3pV63MQPKpOJv649MSsCsabbCMWqqwheTPHm2qlTLmK8ynIAl7Fpfzv5oz3mJiGxPT5Aabzt9U9Tx6TlC/vk2cxj+kF7vLG5A4vjMmafUVAq10SQrGJ63iTLBIGFinbFN7M1txqkYmltYhdniZVnQDaQN6QMMlbUDEYvbch6D2WHjuGjNDOav6kgho12PuoMSw5P25dU7ELBNMNiaWEqasn/5K/7slh9XDnLg59INldbEfC94DYNi3b9EXkBdTJGTwdEM8To21iNDglFnQAC4YDU+vXIn3qSrYG2YRhGC9NA1tTSl9CHTWuhEVo8OS269KzaHDY2boKd4ciG3mfmmzc0WR+BzPiDodmTeQ2oOm8X2y6Xj32bZU6KPaczB80Wj9I9N8ncwi1ajz/bDqtWb8nH77Gq+zN3gE="
	mxGraphModel, err := libdrawio.Decompress(tests)
	if err != nil {
		t.Fail()
	}

	ids := mxGraphModel.GetIds()

	if ids[0] != "0" {
		t.Fail()
	}

	if ids[1] != "1" {
		t.Fail()
	}

	if ids[2] != "eCHvXLgnRnpgWLnYiHm2-1" {
		t.Fail()
	}

	if ids[3] != "ap1ao7U65gjy1zXMJ0Sz-1" {
		t.Fail()
	}

	if ids[4] != "ap1ao7U65gjy1zXMJ0Sz-3" {
		t.Fail()
	}

	if ids[5] != "ap1ao7U65gjy1zXMJ0Sz-2" {
		t.Fail()
	}

	if ids[6] != "ap1ao7U65gjy1zXMJ0Sz-4" {
		t.Fail()
	}
}

func TestGetIdsFromEmptyMxGraphModel(t *testing.T) {
	tests := &libdrawio.MxGraphModel{}

	ids := tests.GetIds()
	if len(ids) != 0 {
		t.Fail()
	}
}

func TestFindObject(t *testing.T) {
	tests := "xZVRT4MwEMc/DY8mtExkr5vTxeiLi9l87MZZagolXSdsn94ixwDHFIzGJ3r/O9rr7+5Sx5vG+a1mafSgQpAOdcPc8a4dSglxffsplH2pBCQoBa5FiEG1sBAHQNFFdSdC2LYCjVLSiLQtblSSwMa0NKa1ytphL0q2T00ZhxNhsWHyVF2K0ER4C3pV63MQPKpOJv649MSsCsabbCMWqqwheTPHm2qlTLmK8ynIAl7Fpfzv5oz3mJiGxPT5Aabzt9U9Tx6TlC/vk2cxj+kF7vLG5A4vjMmafUVAq10SQrGJ63iTLBIGFinbFN7M1txqkYmltYhdniZVnQDaQN6QMMlbUDEYvbch6D2WHjuGjNDOav6kgho12PuoMSw5P25dU7ELBNMNiaWEqasn/5K/7slh9XDnLg59INldbEfC94DYNi3b9EXkBdTJGTwdEM8To21iNDglFnQAC4YDU+vXIn3qSrYG2YRhGC9NA1tTSl9CHTWuhEVo8OS269KzaHDY2boKd4ciG3mfmmzc0WR+BzPiDodmTeQ2oOm8X2y6Xj32bZU6KPaczB80Wj9I9N8ncwi1ajz/bDqtWb8nH77Gq+zN3gE="
	mxGraphModel, err := libdrawio.Decompress(tests)
	if err != nil {
		t.Fail()
	}

	obj := mxGraphModel.FindObject("ap1ao7U65gjy1zXMJ0Sz-4")
	if obj.Id != "ap1ao7U65gjy1zXMJ0Sz-4" {
		t.Fail()
	}
}

func TestFindObjectNotFound(t *testing.T) {
	tests := "xZVRT4MwEMc/DY8mtExkr5vTxeiLi9l87MZZagolXSdsn94ixwDHFIzGJ3r/O9rr7+5Sx5vG+a1mafSgQpAOdcPc8a4dSglxffsplH2pBCQoBa5FiEG1sBAHQNFFdSdC2LYCjVLSiLQtblSSwMa0NKa1ytphL0q2T00ZhxNhsWHyVF2K0ER4C3pV63MQPKpOJv649MSsCsabbCMWqqwheTPHm2qlTLmK8ynIAl7Fpfzv5oz3mJiGxPT5Aabzt9U9Tx6TlC/vk2cxj+kF7vLG5A4vjMmafUVAq10SQrGJ63iTLBIGFinbFN7M1txqkYmltYhdniZVnQDaQN6QMMlbUDEYvbch6D2WHjuGjNDOav6kgho12PuoMSw5P25dU7ELBNMNiaWEqasn/5K/7slh9XDnLg59INldbEfC94DYNi3b9EXkBdTJGTwdEM8To21iNDglFnQAC4YDU+vXIn3qSrYG2YRhGC9NA1tTSl9CHTWuhEVo8OS269KzaHDY2boKd4ciG3mfmmzc0WR+BzPiDodmTeQ2oOm8X2y6Xj32bZU6KPaczB80Wj9I9N8ncwi1ajz/bDqtWb8nH77Gq+zN3gE="
	mxGraphModel, err := libdrawio.Decompress(tests)
	if err != nil {
		t.Fail()
	}

	obj := mxGraphModel.FindObject("aaa")
	if obj != nil {
		t.Fail()
	}
}

func TestFindObjectFromEmptyMxGraphModel(t *testing.T) {
	tests := &libdrawio.MxGraphModel{}
	obj := tests.FindObject("aaa")

	if obj != nil {
		t.Fail()
	}
}

func TestFindMxCell(t *testing.T) {
	tests := "xZVRT4MwEMc/DY8mtExkr5vTxeiLi9l87MZZagolXSdsn94ixwDHFIzGJ3r/O9rr7+5Sx5vG+a1mafSgQpAOdcPc8a4dSglxffsplH2pBCQoBa5FiEG1sBAHQNFFdSdC2LYCjVLSiLQtblSSwMa0NKa1ytphL0q2T00ZhxNhsWHyVF2K0ER4C3pV63MQPKpOJv649MSsCsabbCMWqqwheTPHm2qlTLmK8ynIAl7Fpfzv5oz3mJiGxPT5Aabzt9U9Tx6TlC/vk2cxj+kF7vLG5A4vjMmafUVAq10SQrGJ63iTLBIGFinbFN7M1txqkYmltYhdniZVnQDaQN6QMMlbUDEYvbch6D2WHjuGjNDOav6kgho12PuoMSw5P25dU7ELBNMNiaWEqasn/5K/7slh9XDnLg59INldbEfC94DYNi3b9EXkBdTJGTwdEM8To21iNDglFnQAC4YDU+vXIn3qSrYG2YRhGC9NA1tTSl9CHTWuhEVo8OS269KzaHDY2boKd4ciG3mfmmzc0WR+BzPiDodmTeQ2oOm8X2y6Xj32bZU6KPaczB80Wj9I9N8ncwi1ajz/bDqtWb8nH77Gq+zN3gE="
	mxGraphModel, err := libdrawio.Decompress(tests)
	if err != nil {
		t.Fail()
	}

	obj := mxGraphModel.FindMxCell("ap1ao7U65gjy1zXMJ0Sz-2")
	if obj.Id != "ap1ao7U65gjy1zXMJ0Sz-2" || obj.Parent != "ap1ao7U65gjy1zXMJ0Sz-4" {
		t.Fail()
	}
}

func TestFindMxCellNotFound(t *testing.T) {
	tests := "xZVRT4MwEMc/DY8mtExkr5vTxeiLi9l87MZZagolXSdsn94ixwDHFIzGJ3r/O9rr7+5Sx5vG+a1mafSgQpAOdcPc8a4dSglxffsplH2pBCQoBa5FiEG1sBAHQNFFdSdC2LYCjVLSiLQtblSSwMa0NKa1ytphL0q2T00ZhxNhsWHyVF2K0ER4C3pV63MQPKpOJv649MSsCsabbCMWqqwheTPHm2qlTLmK8ynIAl7Fpfzv5oz3mJiGxPT5Aabzt9U9Tx6TlC/vk2cxj+kF7vLG5A4vjMmafUVAq10SQrGJ63iTLBIGFinbFN7M1txqkYmltYhdniZVnQDaQN6QMMlbUDEYvbch6D2WHjuGjNDOav6kgho12PuoMSw5P25dU7ELBNMNiaWEqasn/5K/7slh9XDnLg59INldbEfC94DYNi3b9EXkBdTJGTwdEM8To21iNDglFnQAC4YDU+vXIn3qSrYG2YRhGC9NA1tTSl9CHTWuhEVo8OS269KzaHDY2boKd4ciG3mfmmzc0WR+BzPiDodmTeQ2oOm8X2y6Xj32bZU6KPaczB80Wj9I9N8ncwi1ajz/bDqtWb8nH77Gq+zN3gE="
	mxGraphModel, err := libdrawio.Decompress(tests)
	if err != nil {
		t.Fail()
	}

	obj := mxGraphModel.FindMxCell("aa")
	if obj != nil {
		t.Fail()
	}
}

func TestFindMxCellFromEmptyMxGraphModel(t *testing.T) {
	tests := &libdrawio.MxGraphModel{}
	obj := tests.FindMxCell("aaa")

	if obj != nil {
		t.Fail()
	}
}

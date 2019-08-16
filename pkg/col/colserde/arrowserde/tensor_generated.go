// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package arrowserde

import flatbuffers "github.com/google/flatbuffers/go"

type SparseTensorIndex = byte
const (
	SparseTensorIndexNONE SparseTensorIndex = 0
	SparseTensorIndexSparseTensorIndexCOO SparseTensorIndex = 1
	SparseTensorIndexSparseMatrixIndexCSR SparseTensorIndex = 2
)

var EnumNamesSparseTensorIndex = map[SparseTensorIndex]string{
	SparseTensorIndexNONE:"NONE",
	SparseTensorIndexSparseTensorIndexCOO:"SparseTensorIndexCOO",
	SparseTensorIndexSparseMatrixIndexCSR:"SparseMatrixIndexCSR",
}

/// ----------------------------------------------------------------------
/// Data structures for dense tensors
/// Shape data for a single axis in a tensor
type TensorDim struct {
	_tab flatbuffers.Table
}

func GetRootAsTensorDim(buf []byte, offset flatbuffers.UOffsetT) *TensorDim {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TensorDim{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *TensorDim) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TensorDim) Table() flatbuffers.Table {
	return rcv._tab
}

/// Length of dimension
func (rcv *TensorDim) Size() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

/// Length of dimension
func (rcv *TensorDim) MutateSize(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

/// Name of the dimension, optional
func (rcv *TensorDim) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Name of the dimension, optional
func TensorDimStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func TensorDimAddSize(builder *flatbuffers.Builder, size int64) {
	builder.PrependInt64Slot(0, size, 0)
}
func TensorDimAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(name), 0)
}
func TensorDimEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type Tensor struct {
	_tab flatbuffers.Table
}

func GetRootAsTensor(buf []byte, offset flatbuffers.UOffsetT) *Tensor {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Tensor{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Tensor) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Tensor) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Tensor) TypeType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Tensor) MutateTypeType(n byte) bool {
	return rcv._tab.MutateByteSlot(4, n)
}

/// The type of data contained in a value cell. Currently only fixed-width
/// value types are supported, no strings or nested types
func (rcv *Tensor) Type(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

/// The type of data contained in a value cell. Currently only fixed-width
/// value types are supported, no strings or nested types
/// The dimensions of the tensor, optionally named
func (rcv *Tensor) Shape(obj *TensorDim, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Tensor) ShapeLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// The dimensions of the tensor, optionally named
/// Non-negative byte offsets to advance one value cell along each dimension
func (rcv *Tensor) Strides(j int) int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *Tensor) StridesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// Non-negative byte offsets to advance one value cell along each dimension
/// The location and size of the tensor's data
func (rcv *Tensor) Data(obj *Buffer) *Buffer {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Buffer)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// The location and size of the tensor's data
func TensorStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func TensorAddTypeType(builder *flatbuffers.Builder, typeType byte) {
	builder.PrependByteSlot(0, typeType, 0)
}
func TensorAddType(builder *flatbuffers.Builder, type_ flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(type_), 0)
}
func TensorAddShape(builder *flatbuffers.Builder, shape flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(shape), 0)
}
func TensorStartShapeVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TensorAddStrides(builder *flatbuffers.Builder, strides flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(strides), 0)
}
func TensorStartStridesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func TensorAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependStructSlot(4, flatbuffers.UOffsetT(data), 0)
}
func TensorEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
/// ----------------------------------------------------------------------
/// EXPERIMENTAL: Data structures for sparse tensors
/// Coodinate format of sparse tensor index.
type SparseTensorIndexCOO struct {
	_tab flatbuffers.Table
}

func GetRootAsSparseTensorIndexCOO(buf []byte, offset flatbuffers.UOffsetT) *SparseTensorIndexCOO {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SparseTensorIndexCOO{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *SparseTensorIndexCOO) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SparseTensorIndexCOO) Table() flatbuffers.Table {
	return rcv._tab
}

/// COO's index list are represented as a NxM matrix,
/// where N is the number of non-zero values,
/// and M is the number of dimensions of a sparse tensor.
/// indicesBuffer stores the location and size of this index matrix.
/// The type of index value is long, so the stride for the index matrix is unnecessary.
///
/// For example, let X be a 2x3x4x5 tensor, and it has the following 6 non-zero values:
///
///   X[0, 1, 2, 0] := 1
///   X[1, 1, 2, 3] := 2
///   X[0, 2, 1, 0] := 3
///   X[0, 1, 3, 0] := 4
///   X[0, 1, 2, 1] := 5
///   X[1, 2, 0, 4] := 6
///
/// In COO format, the index matrix of X is the following 4x6 matrix:
///
///   [[0, 0, 0, 0, 1, 1],
///    [1, 1, 1, 2, 1, 2],
///    [2, 2, 3, 1, 2, 0],
///    [0, 1, 0, 0, 3, 4]]
///
/// Note that the indices are sorted in lexcographical order.
func (rcv *SparseTensorIndexCOO) IndicesBuffer(obj *Buffer) *Buffer {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Buffer)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// COO's index list are represented as a NxM matrix,
/// where N is the number of non-zero values,
/// and M is the number of dimensions of a sparse tensor.
/// indicesBuffer stores the location and size of this index matrix.
/// The type of index value is long, so the stride for the index matrix is unnecessary.
///
/// For example, let X be a 2x3x4x5 tensor, and it has the following 6 non-zero values:
///
///   X[0, 1, 2, 0] := 1
///   X[1, 1, 2, 3] := 2
///   X[0, 2, 1, 0] := 3
///   X[0, 1, 3, 0] := 4
///   X[0, 1, 2, 1] := 5
///   X[1, 2, 0, 4] := 6
///
/// In COO format, the index matrix of X is the following 4x6 matrix:
///
///   [[0, 0, 0, 0, 1, 1],
///    [1, 1, 1, 2, 1, 2],
///    [2, 2, 3, 1, 2, 0],
///    [0, 1, 0, 0, 3, 4]]
///
/// Note that the indices are sorted in lexcographical order.
func SparseTensorIndexCOOStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func SparseTensorIndexCOOAddIndicesBuffer(
	builder *flatbuffers.Builder, indicesBuffer flatbuffers.UOffsetT,
) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(indicesBuffer), 0)
}
func SparseTensorIndexCOOEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
/// Compressed Sparse Row format, that is matrix-specific.
type SparseMatrixIndexCSR struct {
	_tab flatbuffers.Table
}

func GetRootAsSparseMatrixIndexCSR(buf []byte, offset flatbuffers.UOffsetT) *SparseMatrixIndexCSR {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SparseMatrixIndexCSR{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *SparseMatrixIndexCSR) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SparseMatrixIndexCSR) Table() flatbuffers.Table {
	return rcv._tab
}

/// indptrBuffer stores the location and size of indptr array that
/// represents the range of the rows.
/// The i-th row spans from indptr[i] to indptr[i+1] in the data.
/// The length of this array is 1 + (the number of rows), and the type
/// of index value is long.
///
/// For example, let X be the following 6x4 matrix:
///
///   X := [[0, 1, 2, 0],
///         [0, 0, 3, 0],
///         [0, 4, 0, 5],
///         [0, 0, 0, 0],
///         [6, 0, 7, 8],
///         [0, 9, 0, 0]].
///
/// The array of non-zero values in X is:
///
///   values(X) = [1, 2, 3, 4, 5, 6, 7, 8, 9].
///
/// And the indptr of X is:
///
///   indptr(X) = [0, 2, 3, 5, 5, 8, 10].
func (rcv *SparseMatrixIndexCSR) IndptrBuffer(obj *Buffer) *Buffer {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Buffer)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// indptrBuffer stores the location and size of indptr array that
/// represents the range of the rows.
/// The i-th row spans from indptr[i] to indptr[i+1] in the data.
/// The length of this array is 1 + (the number of rows), and the type
/// of index value is long.
///
/// For example, let X be the following 6x4 matrix:
///
///   X := [[0, 1, 2, 0],
///         [0, 0, 3, 0],
///         [0, 4, 0, 5],
///         [0, 0, 0, 0],
///         [6, 0, 7, 8],
///         [0, 9, 0, 0]].
///
/// The array of non-zero values in X is:
///
///   values(X) = [1, 2, 3, 4, 5, 6, 7, 8, 9].
///
/// And the indptr of X is:
///
///   indptr(X) = [0, 2, 3, 5, 5, 8, 10].
/// indicesBuffer stores the location and size of the array that
/// contains the column indices of the corresponding non-zero values.
/// The type of index value is long.
///
/// For example, the indices of the above X is:
///
///   indices(X) = [1, 2, 2, 1, 3, 0, 2, 3, 1].
func (rcv *SparseMatrixIndexCSR) IndicesBuffer(obj *Buffer) *Buffer {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Buffer)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// indicesBuffer stores the location and size of the array that
/// contains the column indices of the corresponding non-zero values.
/// The type of index value is long.
///
/// For example, the indices of the above X is:
///
///   indices(X) = [1, 2, 2, 1, 3, 0, 2, 3, 1].
func SparseMatrixIndexCSRStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func SparseMatrixIndexCSRAddIndptrBuffer(
	builder *flatbuffers.Builder, indptrBuffer flatbuffers.UOffsetT,
) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(indptrBuffer), 0)
}
func SparseMatrixIndexCSRAddIndicesBuffer(
	builder *flatbuffers.Builder, indicesBuffer flatbuffers.UOffsetT,
) {
	builder.PrependStructSlot(1, flatbuffers.UOffsetT(indicesBuffer), 0)
}
func SparseMatrixIndexCSREnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
type SparseTensor struct {
	_tab flatbuffers.Table
}

func GetRootAsSparseTensor(buf []byte, offset flatbuffers.UOffsetT) *SparseTensor {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SparseTensor{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *SparseTensor) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SparseTensor) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SparseTensor) TypeType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SparseTensor) MutateTypeType(n byte) bool {
	return rcv._tab.MutateByteSlot(4, n)
}

/// The type of data contained in a value cell.
/// Currently only fixed-width value types are supported,
/// no strings or nested types.
func (rcv *SparseTensor) Type(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

/// The type of data contained in a value cell.
/// Currently only fixed-width value types are supported,
/// no strings or nested types.
/// The dimensions of the tensor, optionally named.
func (rcv *SparseTensor) Shape(obj *TensorDim, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *SparseTensor) ShapeLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// The dimensions of the tensor, optionally named.
/// The number of non-zero values in a sparse tensor.
func (rcv *SparseTensor) NonZeroLength() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

/// The number of non-zero values in a sparse tensor.
func (rcv *SparseTensor) MutateNonZeroLength(n int64) bool {
	return rcv._tab.MutateInt64Slot(10, n)
}

func (rcv *SparseTensor) SparseIndexType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SparseTensor) MutateSparseIndexType(n byte) bool {
	return rcv._tab.MutateByteSlot(12, n)
}

/// Sparse tensor index
func (rcv *SparseTensor) SparseIndex(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

/// Sparse tensor index
/// The location and size of the tensor's data
func (rcv *SparseTensor) Data(obj *Buffer) *Buffer {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Buffer)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// The location and size of the tensor's data
func SparseTensorStart(builder *flatbuffers.Builder) {
	builder.StartObject(7)
}
func SparseTensorAddTypeType(builder *flatbuffers.Builder, typeType byte) {
	builder.PrependByteSlot(0, typeType, 0)
}
func SparseTensorAddType(builder *flatbuffers.Builder, type_ flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(type_), 0)
}
func SparseTensorAddShape(builder *flatbuffers.Builder, shape flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(shape), 0)
}
func SparseTensorStartShapeVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func SparseTensorAddNonZeroLength(builder *flatbuffers.Builder, nonZeroLength int64) {
	builder.PrependInt64Slot(3, nonZeroLength, 0)
}
func SparseTensorAddSparseIndexType(builder *flatbuffers.Builder, sparseIndexType byte) {
	builder.PrependByteSlot(4, sparseIndexType, 0)
}
func SparseTensorAddSparseIndex(builder *flatbuffers.Builder, sparseIndex flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(sparseIndex), 0)
}
func SparseTensorAddData(builder *flatbuffers.Builder, data flatbuffers.UOffsetT) {
	builder.PrependStructSlot(6, flatbuffers.UOffsetT(data), 0)
}
func SparseTensorEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
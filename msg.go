/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"io"
	"math"
)

type Msg struct {
	data     []byte
	dataSize int
}

func BuildMsgWithSize(dataSize int) *Msg {
	return &Msg{make([]byte, 0, dataSize), 0}
}

func BuildMsg(data []byte) *Msg {
	return &Msg{data, 0}
}

func (Msg *Msg) BuildMsgWithCap(dataSize int) *Msg {

	dm_build_371 := len(Msg.data)
	dm_build_372 := cap(Msg.data)

	if dm_build_371+dataSize <= dm_build_372 {
		Msg.data = Msg.data[:dm_build_371+dataSize]
	} else {

		var calCap = int64(math.Max(float64(2*dm_build_372), float64(dataSize+dm_build_371)))

		nbuf := make([]byte, dataSize+dm_build_371, calCap)
		copy(nbuf, Msg.data)
		Msg.data = nbuf
	}

	return Msg
}

func (Msg *Msg) Dm_build_373() int {
	return len(Msg.data)
}

func (Msg *Msg) Dm_build_375(dm_build_377 int) *Msg {
	for i := dm_build_377; i < len(Msg.data); i++ {
		Msg.data[i] = 0
	}
	Msg.data = Msg.data[:dm_build_377]
	return Msg
}

func (Msg *Msg) Dm_build_378(dm_build_380 int) *Msg {
	Msg.dataSize = dm_build_380
	return Msg
}

func (Msg *Msg) Dm_build_381() int {
	return Msg.dataSize
}

func (Msg *Msg) Dm_build_383(dm_build_385 bool) int {
	return len(Msg.data) - Msg.dataSize
}

func (Msg *Msg) Dm_build_386(dm_build_388 int, dm_build_389 bool, dm_build_390 bool) *Msg {

	if dm_build_389 {
		if dm_build_390 {
			Msg.BuildMsgWithCap(dm_build_388)
		} else {
			Msg.data = Msg.data[:len(Msg.data)-dm_build_388]
		}
	} else {
		if dm_build_390 {
			Msg.dataSize += dm_build_388
		} else {
			Msg.dataSize -= dm_build_388
		}
	}

	return Msg
}

func (Msg *Msg) Dm_build_391(dm_build_393 io.Reader, dm_build_394 int) (int, error) {
	dm_build_395 := len(Msg.data)
	Msg.BuildMsgWithCap(dm_build_394)
	dm_build_396 := 0
	for dm_build_394 > 0 {
		n, err := dm_build_393.Read(Msg.data[dm_build_395+dm_build_396:])
		if n > 0 && err == io.EOF {
			dm_build_396 += n
			Msg.data = Msg.data[:dm_build_395+dm_build_396]
			return dm_build_396, nil
		} else if n > 0 && err == nil {
			dm_build_394 -= n
			dm_build_396 += n
		} else if n == 0 && err != nil {
			return -1, ECGO_COMMUNITION_ERROR.addDetailln(err.Error()).throw()
		}
	}

	return dm_build_396, nil
}

func (Msg *Msg) Dm_build_397(dm_build_399 io.Writer) (*Msg, error) {
	if _, err := dm_build_399.Write(Msg.data); err != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetailln(err.Error()).throw()
	}
	return Msg, nil
}

func (Msg *Msg) Dm_build_400(dm_build_402 bool) int {
	dm_build_403 := len(Msg.data)
	Msg.BuildMsgWithCap(1)

	if dm_build_402 {
		return copy(Msg.data[dm_build_403:], []byte{1})
	} else {
		return copy(Msg.data[dm_build_403:], []byte{0})
	}
}

func (Msg *Msg) Dm_build_404(dm_build_406 byte) int {
	dm_build_407 := len(Msg.data)
	Msg.BuildMsgWithCap(1)

	return copy(Msg.data[dm_build_407:], Packet.Dm_build_179(dm_build_406))
}

func (Msg *Msg) Dm_build_408(dm_build_410 int8) int {
	dm_build_411 := len(Msg.data)
	Msg.BuildMsgWithCap(1)

	return copy(Msg.data[dm_build_411:], Packet.Dm_build_182(dm_build_410))
}

func (dm_build_413 *Msg) Dm_build_412(dm_build_414 int16) int {
	dm_build_415 := len(dm_build_413.data)
	dm_build_413.BuildMsgWithCap(2)

	return copy(dm_build_413.data[dm_build_415:], Packet.Dm_build_185(dm_build_414))
}

func (dm_build_417 *Msg) Dm_build_416(dm_build_418 int32) int {
	dm_build_419 := len(dm_build_417.data)
	dm_build_417.BuildMsgWithCap(4)

	return copy(dm_build_417.data[dm_build_419:], Packet.Dm_build_188(dm_build_418))
}

func (dm_build_421 *Msg) Dm_build_420(dm_build_422 uint8) int {
	dm_build_423 := len(dm_build_421.data)
	dm_build_421.BuildMsgWithCap(1)

	return copy(dm_build_421.data[dm_build_423:], Packet.Dm_build_200(dm_build_422))
}

func (dm_build_425 *Msg) Dm_build_424(dm_build_426 uint16) int {
	dm_build_427 := len(dm_build_425.data)
	dm_build_425.BuildMsgWithCap(2)

	return copy(dm_build_425.data[dm_build_427:], Packet.Dm_build_203(dm_build_426))
}

func (dm_build_429 *Msg) Dm_build_428(dm_build_430 uint32) int {
	dm_build_431 := len(dm_build_429.data)
	dm_build_429.BuildMsgWithCap(4)

	return copy(dm_build_429.data[dm_build_431:], Packet.Dm_build_206(dm_build_430))
}

func (dm_build_433 *Msg) Dm_build_432(dm_build_434 uint64) int {
	dm_build_435 := len(dm_build_433.data)
	dm_build_433.BuildMsgWithCap(8)

	return copy(dm_build_433.data[dm_build_435:], Packet.Dm_build_209(dm_build_434))
}

func (dm_build_437 *Msg) Dm_build_436(dm_build_438 float32) int {
	dm_build_439 := len(dm_build_437.data)
	dm_build_437.BuildMsgWithCap(4)

	return copy(dm_build_437.data[dm_build_439:], Packet.Dm_build_206(math.Float32bits(dm_build_438)))
}

func (dm_build_441 *Msg) Dm_build_440(dm_build_442 float64) int {
	dm_build_443 := len(dm_build_441.data)
	dm_build_441.BuildMsgWithCap(8)

	return copy(dm_build_441.data[dm_build_443:], Packet.Dm_build_209(math.Float64bits(dm_build_442)))
}

func (dm_build_445 *Msg) Dm_build_444(dm_build_446 []byte) int {
	dm_build_447 := len(dm_build_445.data)
	dm_build_445.BuildMsgWithCap(len(dm_build_446))
	return copy(dm_build_445.data[dm_build_447:], dm_build_446)
}

func (dm_build_449 *Msg) Dm_build_448(dm_build_450 []byte) int {
	return dm_build_449.Dm_build_416(int32(len(dm_build_450))) + dm_build_449.Dm_build_444(dm_build_450)
}

func (dm_build_452 *Msg) Dm_build_451(dm_build_453 []byte) int {
	return dm_build_452.Dm_build_420(uint8(len(dm_build_453))) + dm_build_452.Dm_build_444(dm_build_453)
}

func (dm_build_455 *Msg) Dm_build_454(dm_build_456 []byte) int {
	return dm_build_455.Dm_build_424(uint16(len(dm_build_456))) + dm_build_455.Dm_build_444(dm_build_456)
}

func (dm_build_458 *Msg) Dm_build_457(dm_build_459 []byte) int {
	return dm_build_458.Dm_build_444(dm_build_459) + dm_build_458.Dm_build_404(0)
}

func (dm_build_461 *Msg) Dm_build_460(dm_build_462 string, dm_build_463 string, dm_build_464 *DmConnection) int {
	dm_build_465 := Packet.Dm_build_217(dm_build_462, dm_build_463, dm_build_464)
	return dm_build_461.Dm_build_448(dm_build_465)
}

func (dm_build_467 *Msg) Dm_build_466(dm_build_468 string, dm_build_469 string, dm_build_470 *DmConnection) int {
	dm_build_471 := Packet.Dm_build_217(dm_build_468, dm_build_469, dm_build_470)
	return dm_build_467.Dm_build_451(dm_build_471)
}

func (dm_build_473 *Msg) Dm_build_472(dm_build_474 string, dm_build_475 string, dm_build_476 *DmConnection) int {
	dm_build_477 := Packet.Dm_build_217(dm_build_474, dm_build_475, dm_build_476)
	return dm_build_473.Dm_build_454(dm_build_477)
}

func (dm_build_479 *Msg) Dm_build_478(dm_build_480 string, dm_build_481 string, dm_build_482 *DmConnection) int {
	dm_build_483 := Packet.Dm_build_217(dm_build_480, dm_build_481, dm_build_482)
	return dm_build_479.Dm_build_457(dm_build_483)
}

func (dm_build_485 *Msg) Dm_build_484() byte {
	dm_build_486 := Packet.Dm_build_94(dm_build_485.data, dm_build_485.dataSize)
	dm_build_485.dataSize++
	return dm_build_486
}

func (dm_build_488 *Msg) Dm_build_487() int16 {
	dm_build_489 := Packet.Dm_build_98(dm_build_488.data, dm_build_488.dataSize)
	dm_build_488.dataSize += 2
	return dm_build_489
}

func (dm_build_491 *Msg) Dm_build_490() int32 {
	dm_build_492 := Packet.Dm_build_103(dm_build_491.data, dm_build_491.dataSize)
	dm_build_491.dataSize += 4
	return dm_build_492
}

func (dm_build_494 *Msg) Dm_build_493() int64 {
	dm_build_495 := Packet.Dm_build_108(dm_build_494.data, dm_build_494.dataSize)
	dm_build_494.dataSize += 8
	return dm_build_495
}

func (dm_build_497 *Msg) Dm_build_496() float32 {
	dm_build_498 := Packet.Dm_build_113(dm_build_497.data, dm_build_497.dataSize)
	dm_build_497.dataSize += 4
	return dm_build_498
}

func (dm_build_500 *Msg) Dm_build_499() float64 {
	dm_build_501 := Packet.Dm_build_117(dm_build_500.data, dm_build_500.dataSize)
	dm_build_500.dataSize += 8
	return dm_build_501
}

func (dm_build_503 *Msg) Dm_build_502() uint8 {
	dm_build_504 := Packet.Dm_build_121(dm_build_503.data, dm_build_503.dataSize)
	dm_build_503.dataSize += 1
	return dm_build_504
}

func (dm_build_506 *Msg) Dm_build_505() uint16 {
	dm_build_507 := Packet.Dm_build_125(dm_build_506.data, dm_build_506.dataSize)
	dm_build_506.dataSize += 2
	return dm_build_507
}

func (dm_build_509 *Msg) Dm_build_508() uint32 {
	dm_build_510 := Packet.Dm_build_130(dm_build_509.data, dm_build_509.dataSize)
	dm_build_509.dataSize += 4
	return dm_build_510
}

func (dm_build_512 *Msg) Dm_build_511(dm_build_513 int) []byte {
	dm_build_514 := Packet.Dm_build_152(dm_build_512.data, dm_build_512.dataSize, dm_build_513)
	dm_build_512.dataSize += dm_build_513
	return dm_build_514
}

func (dm_build_516 *Msg) Dm_build_515() []byte {
	return dm_build_516.Dm_build_511(int(dm_build_516.Dm_build_490()))
}

func (dm_build_518 *Msg) Dm_build_517() []byte {
	return dm_build_518.Dm_build_511(int(dm_build_518.Dm_build_484()))
}

func (dm_build_520 *Msg) Dm_build_519() []byte {
	return dm_build_520.Dm_build_511(int(dm_build_520.Dm_build_487()))
}

func (dm_build_522 *Msg) Dm_build_521(dm_build_523 int) []byte {
	return dm_build_522.Dm_build_511(dm_build_523)
}

func (dm_build_525 *Msg) Dm_build_524() []byte {
	dm_build_526 := 0
	for dm_build_525.Dm_build_484() != 0 {
		dm_build_526++
	}
	dm_build_525.Dm_build_386(dm_build_526, false, false)
	return dm_build_525.Dm_build_511(dm_build_526)
}

func (dm_build_528 *Msg) Dm_build_527(dm_build_529 int, dm_build_530 string, dm_build_531 *DmConnection) string {
	return Packet.Dm_build_254(dm_build_528.Dm_build_511(dm_build_529), dm_build_530, dm_build_531)
}

func (dm_build_533 *Msg) Dm_build_532(dm_build_534 string, dm_build_535 *DmConnection) string {
	return Packet.Dm_build_254(dm_build_533.Dm_build_515(), dm_build_534, dm_build_535)
}

func (dm_build_537 *Msg) Dm_build_536(dm_build_538 string, dm_build_539 *DmConnection) string {
	return Packet.Dm_build_254(dm_build_537.Dm_build_517(), dm_build_538, dm_build_539)
}

func (dm_build_541 *Msg) Dm_build_540(dm_build_542 string, dm_build_543 *DmConnection) string {
	return Packet.Dm_build_254(dm_build_541.Dm_build_519(), dm_build_542, dm_build_543)
}

func (dm_build_545 *Msg) Dm_build_544(dm_build_546 string, dm_build_547 *DmConnection) string {
	return Packet.Dm_build_254(dm_build_545.Dm_build_524(), dm_build_546, dm_build_547)
}

func (dm_build_549 *Msg) Dm_build_548(dm_build_550 int, dm_build_551 byte) int {
	return dm_build_549.Dm_build_584(dm_build_550, Packet.Dm_build_179(dm_build_551))
}

func (dm_build_553 *Msg) Dm_build_552(dm_build_554 int, dm_build_555 int16) int {
	return dm_build_553.Dm_build_584(dm_build_554, Packet.Dm_build_185(dm_build_555))
}

func (dm_build_557 *Msg) Dm_build_556(dm_build_558 int, dm_build_559 int32) int {
	return dm_build_557.Dm_build_584(dm_build_558, Packet.Dm_build_188(dm_build_559))
}

func (dm_build_561 *Msg) Dm_build_560(dm_build_562 int, dm_build_563 int64) int {
	return dm_build_561.Dm_build_584(dm_build_562, Packet.Dm_build_191(dm_build_563))
}

func (dm_build_565 *Msg) Dm_build_564(dm_build_566 int, dm_build_567 float32) int {
	return dm_build_565.Dm_build_584(dm_build_566, Packet.Dm_build_194(dm_build_567))
}

func (dm_build_569 *Msg) Dm_build_568(dm_build_570 int, dm_build_571 float64) int {
	return dm_build_569.Dm_build_584(dm_build_570, Packet.Dm_build_197(dm_build_571))
}

func (dm_build_573 *Msg) Dm_build_572(dm_build_574 int, dm_build_575 uint8) int {
	return dm_build_573.Dm_build_584(dm_build_574, Packet.Dm_build_200(dm_build_575))
}

func (dm_build_577 *Msg) Dm_build_576(dm_build_578 int, dm_build_579 uint16) int {
	return dm_build_577.Dm_build_584(dm_build_578, Packet.Dm_build_203(dm_build_579))
}

func (dm_build_581 *Msg) Dm_build_580(dm_build_582 int, dm_build_583 uint32) int {
	return dm_build_581.Dm_build_584(dm_build_582, Packet.Dm_build_206(dm_build_583))
}

func (dm_build_585 *Msg) Dm_build_584(dm_build_586 int, dm_build_587 []byte) int {
	return copy(dm_build_585.data[dm_build_586:], dm_build_587)
}

func (dm_build_589 *Msg) Dm_build_588(dm_build_590 int, dm_build_591 []byte) int {
	return dm_build_589.Dm_build_556(dm_build_590, int32(len(dm_build_591))) + dm_build_589.Dm_build_584(dm_build_590+4, dm_build_591)
}

func (dm_build_593 *Msg) Dm_build_592(dm_build_594 int, dm_build_595 []byte) int {
	return dm_build_593.Dm_build_548(dm_build_594, byte(len(dm_build_595))) + dm_build_593.Dm_build_584(dm_build_594+1, dm_build_595)
}

func (dm_build_597 *Msg) Dm_build_596(dm_build_598 int, dm_build_599 []byte) int {
	return dm_build_597.Dm_build_552(dm_build_598, int16(len(dm_build_599))) + dm_build_597.Dm_build_584(dm_build_598+2, dm_build_599)
}

func (dm_build_601 *Msg) Dm_build_600(dm_build_602 int, dm_build_603 []byte) int {
	return dm_build_601.Dm_build_584(dm_build_602, dm_build_603) + dm_build_601.Dm_build_548(dm_build_602+len(dm_build_603), 0)
}

func (dm_build_605 *Msg) Dm_build_604(dm_build_606 int, dm_build_607 string, dm_build_608 string, dm_build_609 *DmConnection) int {
	return dm_build_605.Dm_build_588(dm_build_606, Packet.Dm_build_217(dm_build_607, dm_build_608, dm_build_609))
}

func (dm_build_611 *Msg) Dm_build_610(dm_build_612 int, dm_build_613 string, dm_build_614 string, dm_build_615 *DmConnection) int {
	return dm_build_611.Dm_build_592(dm_build_612, Packet.Dm_build_217(dm_build_613, dm_build_614, dm_build_615))
}

func (dm_build_617 *Msg) Dm_build_616(dm_build_618 int, dm_build_619 string, dm_build_620 string, dm_build_621 *DmConnection) int {
	return dm_build_617.Dm_build_596(dm_build_618, Packet.Dm_build_217(dm_build_619, dm_build_620, dm_build_621))
}

func (dm_build_623 *Msg) Dm_build_622(dm_build_624 int, dm_build_625 string, dm_build_626 string, dm_build_627 *DmConnection) int {
	return dm_build_623.Dm_build_600(dm_build_624, Packet.Dm_build_217(dm_build_625, dm_build_626, dm_build_627))
}

func (dm_build_629 *Msg) Dm_build_628(dm_build_630 int) byte {
	return Packet.Dm_build_222(dm_build_629.Dm_build_655(dm_build_630, 1))
}

func (dm_build_632 *Msg) Dm_build_631(dm_build_633 int) int16 {
	return Packet.Dm_build_225(dm_build_632.Dm_build_655(dm_build_633, 2))
}

func (dm_build_635 *Msg) Dm_build_634(dm_build_636 int) int32 {
	return Packet.Dm_build_228(dm_build_635.Dm_build_655(dm_build_636, 4))
}

func (dm_build_638 *Msg) Dm_build_637(dm_build_639 int) int64 {
	return Packet.Dm_build_231(dm_build_638.Dm_build_655(dm_build_639, 8))
}

func (dm_build_641 *Msg) Dm_build_640(dm_build_642 int) float32 {
	return Packet.Dm_build_234(dm_build_641.Dm_build_655(dm_build_642, 4))
}

func (dm_build_644 *Msg) Dm_build_643(dm_build_645 int) float64 {
	return Packet.Dm_build_237(dm_build_644.Dm_build_655(dm_build_645, 8))
}

func (dm_build_647 *Msg) Dm_build_646(dm_build_648 int) uint8 {
	return Packet.Dm_build_240(dm_build_647.Dm_build_655(dm_build_648, 1))
}

func (dm_build_650 *Msg) Dm_build_649(dm_build_651 int) uint16 {
	return Packet.Dm_build_243(dm_build_650.Dm_build_655(dm_build_651, 2))
}

func (dm_build_653 *Msg) Dm_build_652(dm_build_654 int) uint32 {
	return Packet.Dm_build_246(dm_build_653.Dm_build_655(dm_build_654, 4))
}

func (dm_build_656 *Msg) Dm_build_655(dm_build_657 int, dm_build_658 int) []byte {
	return dm_build_656.data[dm_build_657 : dm_build_657+dm_build_658]
}

func (dm_build_660 *Msg) Dm_build_659(dm_build_661 int) []byte {
	dm_build_662 := dm_build_660.Dm_build_634(dm_build_661)
	return dm_build_660.Dm_build_655(dm_build_661+4, int(dm_build_662))
}

func (dm_build_664 *Msg) Dm_build_663(dm_build_665 int) []byte {
	dm_build_666 := dm_build_664.Dm_build_628(dm_build_665)
	return dm_build_664.Dm_build_655(dm_build_665+1, int(dm_build_666))
}

func (dm_build_668 *Msg) Dm_build_667(dm_build_669 int) []byte {
	dm_build_670 := dm_build_668.Dm_build_631(dm_build_669)
	return dm_build_668.Dm_build_655(dm_build_669+2, int(dm_build_670))
}

func (dm_build_672 *Msg) Dm_build_671(dm_build_673 int) []byte {
	dm_build_674 := 0
	for dm_build_672.Dm_build_628(dm_build_673) != 0 {
		dm_build_673++
		dm_build_674++
	}

	return dm_build_672.Dm_build_655(dm_build_673-dm_build_674, int(dm_build_674))
}

func (dm_build_676 *Msg) Dm_build_675(dm_build_677 int, dm_build_678 string, dm_build_679 *DmConnection) string {
	return Packet.Dm_build_254(dm_build_676.Dm_build_659(dm_build_677), dm_build_678, dm_build_679)
}

func (dm_build_681 *Msg) Dm_build_680(dm_build_682 int, dm_build_683 string, dm_build_684 *DmConnection) string {
	return Packet.Dm_build_254(dm_build_681.Dm_build_663(dm_build_682), dm_build_683, dm_build_684)
}

func (dm_build_686 *Msg) Dm_build_685(dm_build_687 int, dm_build_688 string, dm_build_689 *DmConnection) string {
	return Packet.Dm_build_254(dm_build_686.Dm_build_667(dm_build_687), dm_build_688, dm_build_689)
}

func (dm_build_691 *Msg) Dm_build_690(dm_build_692 int, dm_build_693 string, dm_build_694 *DmConnection) string {
	return Packet.Dm_build_254(dm_build_691.Dm_build_671(dm_build_692), dm_build_693, dm_build_694)
}

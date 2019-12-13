require './machine'
require "./number"
require "./evaluate"
require "./add"
require "./less_than"
require "./variable"
require "./sequence"
require "./assign"
require "./boolean"
require './do_nothing'


require "test/unit"
require "test/unit/assertions"
Test::Unit::Assertions.use_pp = false

class MachineTest < Test::Unit::TestCase

  def test_add
    machine = Machine.new(Evaluate.new(Add.new(Number.new(1),
                                  Add.new(Number.new(2),
                                          Add.new(Number.new(3),
                                                  Number.new(4))))),
                          {})

    out = capture_output { machine.run }[0]
    expected = <<-eos
eval(1 + 2 + 3 + 4), {}
eval(1 + 2 + 7), {}
eval(1 + 9), {}
eval(10), {}
do-nothing, {}
eos
    assert_equal expected, out
    assert_equal DoNothing.new, machine.statement
  end

  def test_less_than
    machine = Machine.new(Evaluate.new(LessThan.new(Add.new(Number.new(1), Number.new(2)),
                                       Add.new(Number.new(3), Number.new(4)))),
                          {})

    out = capture_output { machine.run }[0]
    expected = <<-eos
eval(1 + 2 < 3 + 4), {}
eval(3 < 3 + 4), {}
eval(3 < 7), {}
eval(true), {}
do-nothing, {}
eos
    assert_equal expected, out
    assert_equal DoNothing.new, machine.statement
  end

  def test_variable
    machine = Machine.new(Evaluate.new(Variable.new(:x)),
                          {x: Number.new(100)})

    out = capture_output { machine.run }[0]
    expected = <<-eos
eval(x), {:x=>«100»}
eval(100), {:x=>«100»}
do-nothing, {:x=>«100»}
eos
    assert_equal expected, out
    assert_equal DoNothing.new, machine.statement
  end

  def test_add_variables
    machine = Machine.new(Evaluate.new(Add.new(Variable.new(:x), Variable.new(:y))),
                          {x: Number.new(100), y: Number.new(200)})

    out = capture_output { machine.run }[0]
    expected = <<-eos
eval(x + y), {:x=>«100», :y=>«200»}
eval(100 + y), {:x=>«100», :y=>«200»}
eval(100 + 200), {:x=>«100», :y=>«200»}
eval(300), {:x=>«100», :y=>«200»}
do-nothing, {:x=>«100», :y=>«200»}
eos
    assert_equal expected, out
    assert_equal DoNothing.new, machine.statement
  end

  def test_assign
    machine = Machine.new(Assign.new(:x, Add.new(Variable.new(:y), Number.new(10))),
                          {y: Number.new(1)})

    out = capture_output { machine.run }[0]
    expected = <<-eos
x = y + 10, {:y=>«1»}
x = 1 + 10, {:y=>«1»}
x = 11, {:y=>«1»}
do-nothing, {:y=>«1», :x=>«11»}
eos
    assert_equal expected, out
    assert_equal DoNothing.new, machine.statement
  end

  def test_sequence_complex
    machine = Machine.new(Sequence.new(Assign.new(:x, Add.new(Number.new(1), Number.new(1))),
                                       Assign.new(:y, Add.new(Variable.new(:x), Number.new(3)))),
                          {})

    out = capture_output { machine.run }[0]
    expected = <<-eos
x = 1 + 1; y = x + 3, {}
x = 2; y = x + 3, {}
do-nothing; y = x + 3, {:x=>«2»}
y = x + 3, {:x=>«2»}
y = 2 + 3, {:x=>«2»}
y = 5, {:x=>«2»}
do-nothing, {:x=>«2», :y=>«5»}
eos
    assert_equal expected, out
    assert_equal DoNothing.new, machine.statement

  end
end

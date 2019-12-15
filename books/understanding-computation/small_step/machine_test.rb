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
require './iff'
require './whilee'
require './multiply'

require './test_setup'

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

  def test_if_true
    machine = Machine.new(If.new(Variable.new(:x),
                                 Assign.new(:y, Number.new(1)),
                                 Assign.new(:y, Number.new(2))),
                          {x: Boolean.new(true)})

    out = capture_output { machine.run }[0]
    expected = <<-eos
if (x) { y = 1 } else { y = 2 }, {:x=>«true»}
if (true) { y = 1 } else { y = 2 }, {:x=>«true»}
y = 1, {:x=>«true»}
do-nothing, {:x=>«true», :y=>«1»}
eos
    assert_equal expected, out
    assert_equal DoNothing.new, machine.statement
  end
  

  def test_if_false
    machine = Machine.new(If.new(Variable.new(:x), Assign.new(:y, Number.new(1)), DoNothing.new),
                          {x: Boolean.new(false)})
    
    out = capture_output { machine.run }[0]
    expected = <<-eos
if (x) { y = 1 } else { do-nothing }, {:x=>«false»}
if (false) { y = 1 } else { do-nothing }, {:x=>«false»}
do-nothing, {:x=>«false»}
eos
    assert_equal expected, out
    assert_equal DoNothing.new, machine.statement
  end


  def test_while
    machine = Machine.new(While.new(LessThan.new(Variable.new(:x), Number.new(5)),
                                    Assign.new(:x, Multiply.new(Variable.new(:x), Number.new(3)))),
                          { x: Number.new(1) })
    
    out = capture_output { machine.run }[0]
    expected = <<-eos
while (x < 5) { x = x * 3 }, {:x=>«1»}
if (x < 5) { x = x * 3; while (x < 5) { x = x * 3 } } else { do-nothing }, {:x=>«1»}
if (1 < 5) { x = x * 3; while (x < 5) { x = x * 3 } } else { do-nothing }, {:x=>«1»}
if (true) { x = x * 3; while (x < 5) { x = x * 3 } } else { do-nothing }, {:x=>«1»}
x = x * 3; while (x < 5) { x = x * 3 }, {:x=>«1»}
x = 1 * 3; while (x < 5) { x = x * 3 }, {:x=>«1»}
x = 3; while (x < 5) { x = x * 3 }, {:x=>«1»}
do-nothing; while (x < 5) { x = x * 3 }, {:x=>«3»}
while (x < 5) { x = x * 3 }, {:x=>«3»}
if (x < 5) { x = x * 3; while (x < 5) { x = x * 3 } } else { do-nothing }, {:x=>«3»}
if (3 < 5) { x = x * 3; while (x < 5) { x = x * 3 } } else { do-nothing }, {:x=>«3»}
if (true) { x = x * 3; while (x < 5) { x = x * 3 } } else { do-nothing }, {:x=>«3»}
x = x * 3; while (x < 5) { x = x * 3 }, {:x=>«3»}
x = 3 * 3; while (x < 5) { x = x * 3 }, {:x=>«3»}
x = 9; while (x < 5) { x = x * 3 }, {:x=>«3»}
do-nothing; while (x < 5) { x = x * 3 }, {:x=>«9»}
while (x < 5) { x = x * 3 }, {:x=>«9»}
if (x < 5) { x = x * 3; while (x < 5) { x = x * 3 } } else { do-nothing }, {:x=>«9»}
if (9 < 5) { x = x * 3; while (x < 5) { x = x * 3 } } else { do-nothing }, {:x=>«9»}
if (false) { x = x * 3; while (x < 5) { x = x * 3 } } else { do-nothing }, {:x=>«9»}
do-nothing, {:x=>«9»}
eos
    assert_equal expected, out
    assert_equal DoNothing.new, machine.statement
  end
  
end

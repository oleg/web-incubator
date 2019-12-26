require './test_setup'

require './concatenate'
require './literal'
require './choose'
require './repeat'

class ConcatenateTest < Test::Unit::TestCase
  
  def test_concatenate_literal
    assert_equal '12', Concatenate.new(Literal.new('1'),
                                       Literal.new('2')).to_s
  end

  def test_concatenate_repeate
    assert_equal '1*2*', Concatenate.new(Repeat.new(Literal.new('1')),
                                         Repeat.new(Literal.new('2'))).to_s
  end
  
  
  def test_concatenate_choose
    assert_equal '(1|2)(3|4)', Concatenate.new(Choose.new(Literal.new('1'), Literal.new('2')),
                                               Choose.new(Literal.new('3'), Literal.new('4'))).to_s
  end

  def test_concatenate_concatenate
    assert_equal 'dogs', Concatenate.new(Concatenate.new(Literal.new('d'), Literal.new('o')),
                                         Concatenate.new(Literal.new('g'), Literal.new('s'))).to_s
  end

  def test_concatenate
    pattern = Concatenate.new(Literal.new('a'), Literal.new('b'))

    assert_false pattern.matches?('a')
    assert_true pattern.matches?('ab')
    assert_false pattern.matches?('abc')
  end

  def test_concatenate_recursive
    pattern = Concatenate.new(Literal.new('a'), Concatenate.new(Literal.new('b'), Literal.new('c')))
    
    assert_false pattern.matches?('a')
    assert_false pattern.matches?('ab')
    assert_true pattern.matches?('abc')
  end
  
end

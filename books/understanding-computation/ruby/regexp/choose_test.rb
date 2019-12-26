require './test_setup'

require './choose'
require './literal'
require './concatenate'
require './repeat'


class ChooseTest < Test::Unit::TestCase
  

  def test_choose_literal
    assert_equal '1|2', Choose.new(Literal.new('1'),
                                   Literal.new('2')).to_s
  end

  def test_choose_repeat
    assert_equal 'a*|b*', Choose.new(Repeat.new(Literal.new('a')),
                                     Repeat.new(Literal.new('b'))).to_s
  end

  def test_choose_concatenate
    assert_equal 'a1|b2', Choose.new(Concatenate.new(Literal.new('a'), Literal.new('1')),
                                     Concatenate.new(Literal.new('b'), Literal.new('2'))).to_s
  end

  def test_choose_choose
    assert_equal 'x|y|z', Choose.new(Choose.new(Literal.new('x'),
                                                Literal.new('y')),
                                     Literal.new('z')).to_s
  end

  def test_matches
    pattern = Choose.new(Literal.new('a'), Literal.new('b'))

    assert_true pattern.matches?('a')
    assert_true pattern.matches?('b')
    assert_false pattern.matches?('c')
  end
  
end

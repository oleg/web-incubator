require './test_setup'

require './repeat'
require './choose'
require './concatenate'
require './literal'


class RepeatTest < Test::Unit::TestCase

  def test_repeate_literal
    assert_equal 'd*', Repeat.new(Literal.new('d')).to_s
  end
  
  def test_repeate_concatenate
    assert_equal '(ab)*', Repeat.new(Concatenate.new(Literal.new('a'), Literal.new('b'))).to_s
  end

  def test_repeate_choose
    assert_equal '(a|b)*', Repeat.new(Choose.new(Literal.new('a'), Literal.new('b'))).to_s
  end

  def test_repeate_repeate
    assert_equal '9**', Repeat.new(Repeat.new(Literal.new('9'))).to_s
  end

  def test_inspect
    assert_equal '/d*/', Repeat.new(Literal.new('d')).inspect
  end

  def test_matches
    pattern = Repeat.new(Literal.new('a'))
    
    assert_true pattern.matches?('')
    assert_true pattern.matches?('a')
    assert_true pattern.matches?('aaaa')
    assert_false pattern.matches?('b')
  end
  
end

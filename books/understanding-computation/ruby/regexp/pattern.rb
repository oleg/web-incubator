module Pattern
  def bracket outer_precedence
    if outer_precedence > precedence
      "(#{to_s})"
    else
      to_s
    end
  end
  
  def inspect
    "/#{self}/"
  end
end

class Empty
  include Pattern

  def precedence
    3
  end
  
  def to_s
    ''
  end
  
end

class Literal < Struct.new(:character)
  include Pattern
  
  def precedence
    3
  end

  def to_s
    character
  end

end

class Repeat < Struct.new(:pattern)
  include Pattern

  def precedence
    2
  end
  
  def to_s
    pattern.bracket(precedence) + '*'
  end

end

class Concatenate < Struct.new(:first, :second)
  include Pattern
  
  def precedence
    1
  end
  
  def to_s
    [first, second].map { |pattern| pattern.bracket(precedence) }.join
  end
  
end

class Choose < Struct.new(:first, :second)
  include Pattern
  
  def precedence
    0
  end
  
  def to_s
    #[first, second].map { |pattern| pattern.bracket(precedence) }.join('|')
    [first, second].join('|')
  end
  
end



require 'test/unit'
require "test/unit/assertions"
Test::Unit::Assertions.use_pp = false


class EmptyTest < Test::Unit::TestCase
  
  def test_to_s
    assert_equal '', Empty.new.to_s
  end

  def test_inspect
    assert_equal '//', Empty.new.inspect
  end

end

class LiteralTest < Test::Unit::TestCase
  
  def test_to_s
    assert_equal 'a', Literal.new('a').to_s
    assert_equal 'b', Literal.new('b').to_s    
  end

  def test_inspect
    assert_equal '/c/', Literal.new('c').inspect
    assert_equal '/d/', Literal.new('d').inspect
  end

end

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

end

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

end

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

end


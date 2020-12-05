require './test_setup'

require './pattern'

require './repeat'
require './concatenate'
require './literal'
require './choose'
require './empty'


class PatternTest < Test::Unit::TestCase

  def setup
    @testcl = Class.new do
      include Pattern
      def precedence; 100 end
      def to_s; "X" end
    end
    
    @testinst = @testcl.new
  end

  def test_bracket_less
    assert_equal "(X)", @testinst.bracket(101)
  end

  def test_bracket_equal
    assert_equal "X", @testinst.bracket(100)
  end

  def test_bracket_greater
    assert_equal "X", @testinst.bracket(1)
  end
  
  def test_inspect
    assert_equal "/X/", @testinst.inspect
  end

end

class PatternMatchesTest < Test::Unit::TestCase

  def test_complex_match
    pattern = Repeat.new(Concatenate.new(Literal.new('a'), Choose.new(Empty.new, Literal.new('b'))))

    assert_equal '/(a(|b))*/', pattern.inspect
    assert_true pattern.matches?('')
    assert_true pattern.matches?('a')
    assert_true pattern.matches?('ab')
    assert_true pattern.matches?('aba')
    assert_true pattern.matches?('abab')
    assert_true pattern.matches?('abaab')
    assert_false pattern.matches?('abba')
  end

end

require './pattern'
require './test_setup'

class PatternTest < Test::Unit::TestCase

  def setup
    @testcl = Class.new do
      include Pattern
      def precedence; 100 end
      def to_s; "X" end
    end
  end

  def test_bracket_less
    testinst = @testcl.new
    assert_equal "(X)", testinst.bracket(101)
  end

  def test_bracket_equal
    testinst = @testcl.new
    assert_equal "X", testinst.bracket(100)
  end

  def test_bracket_greater
    testinst = @testcl.new
    assert_equal "X", testinst.bracket(1)
  end
  
  def test_inspect
    testinst = @testcl.new
    assert_equal "/X/", testinst.inspect
  end

end

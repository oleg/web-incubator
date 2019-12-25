require './literal'
require './test_setup'

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

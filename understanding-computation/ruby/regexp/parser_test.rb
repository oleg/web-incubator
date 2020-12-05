require './test_setup'

require './parser'

class ParserTest < Test::Unit::TestCase

  def test_parser
    pattern = Parser.new.parse('(a(|b))*')
    
    assert_equal '/(a(|b))*/', pattern.inspect
    assert_true pattern.matches?('abaab')
    assert_false pattern.matches?('abba')
  end
  
end

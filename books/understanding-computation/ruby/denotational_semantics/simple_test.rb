require 'treetop'
require './number'
require './while'
require './assign'
require './multiply'
require './variable'
require './add'
require './less_than'

require './test_setup'

class SimpleTest < Test::Unit::TestCase

  def test_simple
    Treetop.load('simple.tt')
    parse_tree = SimpleParser.new.parse('while (x < 5) { x = x * 3 }')
    statement = parse_tree.to_ast
    env_res = eval(statement.to_ruby)[{x: 1}]
    assert_equal Hash[x: 9], env_res
  end
  
end

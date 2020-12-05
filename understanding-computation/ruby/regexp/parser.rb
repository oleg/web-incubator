require 'treetop'

require './concatenate'
require './choose'
require './repeat'
require './literal'
require './empty'


Treetop.load('pattern')

class Parser
  
  def parse pattern
    PatternParser.new.parse(pattern).to_ast
  end
  
end

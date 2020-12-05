require './nfa_rulebook'
require './fa_rule'
require './nfa_design'
require './pattern'

class Literal < Struct.new(:character)
  include Pattern
  
  def precedence
    3
  end

  def to_s
    character
  end

  def to_nfa_design
    start_state = Object.new
    accept_state = Object.new
    
    rulebook = NFARulebook.new([
      FARule.new(start_state, character, accept_state)
    ])
    
    NFADesign.new(start_state, [accept_state], rulebook)
  end
  
end

require './pattern'
require './nfa_rulebook'
require './nfa_design'

class Empty
  include Pattern

  def precedence
    3
  end
  
  def to_s
    ''
  end

  def to_nfa_design
    start_state = Object.new
    accept_states = [start_state]
    rulebook = NFARulebook.new([])

    NFADesign.new(start_state, accept_states, rulebook)
  end
  
end

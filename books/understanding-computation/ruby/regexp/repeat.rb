require './pattern'

class Repeat < Struct.new(:pattern)
  include Pattern

  def precedence
    2
  end
  
  def to_s
    pattern.bracket(precedence) + '*'
  end

  def to_nfa_design
    start_state = Object.new

    d = pattern.to_nfa_design
    
    accept_states = [start_state] + d.accept_states
    
    old_rules = d.rulebook.rules

    repeat_rules = d.accept_states.map { |s| FARule.new(s, nil, d.start_state) }

    connect_rule = FARule.new(start_state, nil, d.start_state)
      
    rulebook = NFARulebook.new(old_rules + repeat_rules + [connect_rule])
    
    NFADesign.new(start_state,
                  accept_states,
                  rulebook)
  end

end

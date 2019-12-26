require './pattern'

class Choose < Struct.new(:first, :second)
  include Pattern
  
  def precedence
    0
  end
  
  def to_s
    #[first, second].map { |pattern| pattern.bracket(precedence) }.join('|')
    [first, second].join('|')
  end

  def to_nfa_design
    start_state = Object.new
    
    fd = first.to_nfa_design
    sd = second.to_nfa_design

    accept_states = fd.accept_states + sd.accept_states

    old_rules = [fd, sd].map{ |d| d.rulebook.rules }.reduce(&:+)
    
    connect_rules = [fd, sd].map { |d| FARule.new(start_state, nil, d.start_state) }
    
    rulebook = NFARulebook.new(old_rules + connect_rules)
    
    NFADesign.new(start_state,
                  accept_states,
                  rulebook)
  end
  
end

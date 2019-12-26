require './pattern'

class Concatenate < Struct.new(:first, :second)
  include Pattern
  
  def precedence
    1
  end
  
  def to_s
    [first, second].map { |pattern| pattern.bracket(precedence) }.join
  end

  def to_nfa_design
    fd = first.to_nfa_design
    sd = second.to_nfa_design
    
    old_rules = [fd, sd].map{ |d| d.rulebook.rules }.reduce(&:+)
    connect_rules = fd.accept_states.map { |s| FARule.new(s, nil, sd.start_state) }
    
    rulebook = NFARulebook.new(old_rules + connect_rules)

    NFADesign.new(fd.start_state,
                  sd.accept_states,
                  rulebook)
  end
  
end

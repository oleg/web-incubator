require './nfa'
require 'set'

class NFADesign < Struct.new(:start_state, :accept_states, :rulebook)

  def accepts? str
    to_nfa.tap { |nfa| nfa.read_string str }.accepting?
  end

  def to_nfa
    NFA.new(Set[start_state], accept_states, rulebook)
  end
  
end

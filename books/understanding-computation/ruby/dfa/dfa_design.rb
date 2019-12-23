class DFADesign < Struct.new(:start_state, :accept_states, :rulebook)
  
  def accepts?(string)
    to_dfa.tap { |dfa| dfa.read_string(string) }.accepting?
  end

  def to_dfa
    DFA.new(start_state, accept_states, rulebook)
  end

end

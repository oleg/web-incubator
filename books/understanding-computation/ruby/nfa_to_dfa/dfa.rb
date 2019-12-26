class DFA < Struct.new(:current_state, :accept_states, :rulebook)
  
  def accepting?
    accept_states.include?(current_state)
  end
  
  def read_character(character)
    self.current_state = rulebook.next_state(self.current_state, character)
  end

  def read_string(str)
    str.chars.each { |ch| read_character(ch) }
  end
  
end


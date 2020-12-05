class NFA < Struct.new(:current_states, :accept_states, :rulebook)
  
  def accepting?
    (current_states & accept_states).any?
  end

  def current_states
    rulebook.follow_free_moves(super)
  end

  def read_character character
    self.current_states = rulebook.next_states(self.current_states, character)
  end

  def read_string str
    str.chars.each { |ch| read_character ch }
  end
  
end

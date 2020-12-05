class NFARulebook < Struct.new(:rules)

  def follow_free_moves states
    more_states = next_states(states, nil)
    
    if more_states.subset?(states)
      states
    else
      follow_free_moves(states + more_states)
    end
  end

  def next_states(states, character)
    states.flat_map { |state| rules_for(state, character) }.map(&:follow).to_set
  end

  def rules_for(state, character)
    rules.select { |rule| rule.applies_to?(state, character)}
  end

  def alphabet
    rules.map(&:character).compact.uniq
  end
  


  def + other
    NFARulebook.new(self.rules + other.rules)
  end

end

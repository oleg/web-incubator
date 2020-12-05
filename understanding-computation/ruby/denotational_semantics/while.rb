class While < Struct.new(:condition, :body)

  def to_ruby
    "-> e { while (#{condition.to_ruby})[e]; e = (#{body.to_ruby})[e]; end; e }"
  end

end

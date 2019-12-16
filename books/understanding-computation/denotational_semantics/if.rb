class If < Struct.new(:condition, :consequence, :alternative)
  
  def to_ruby
    "-> e { 
if (#{condition.to_ruby})[e] 
then (#{consequence.to_ruby})[e] 
else (#{alternative.to_ruby})[e] 
end 
}"
  end
  
end

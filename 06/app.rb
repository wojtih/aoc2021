def read_data
  File.readlines("data.txt").map { |l| l.split(",").map(&:to_i).sort }.first
end

res = Hash.new(0)
read_data.each { |i| res[i] +=1 }

80.times do |day|
  res_new = Hash.new(0)
  res.each do |k,v|
    if k == 0
      res_new[6] += v
      res_new[8] = v
    else
      res_new[k-1] += v  
    end
  end
  res = res_new
end
puts res.values.inject(0) {|i,sum| sum += i}

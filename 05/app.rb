def read_data
  File.readlines("data.txt").map { |l| l.split(" -> ").map {|co| co.split(",").map(&:to_i) }.sort }
end

def squares(x1, y1, x2, y2) 
  len = (y2 - y1).abs
  if y1 == y2
    x1.upto(x2).map { |x| [x, y1] }
  elsif x1 == x2
    if y1 < y2
      (0..len).map {|i| [x1, y1+i]}
    else
      (0..len).map {|i| [x1, y1-i]}
    end
  else
    if y1 <= y2
      (0..len).map {|i| [x1+i, y1+i]}
    else
      (0..len).map {|i| [x1+i, y1-i]}
    end
  end
end

res = Hash.new(0)
read_data.each do |pair|
  squares(*pair.flatten).each do |square|
    key = "#{square[0]}-#{square[1]}"
    res[key] += 1
  end
end

puts res
puts res.count {|_,v| v >1 }

class AoC
# PART 1
  def call
    nums = File.readlines('data.txt').map {|i| i.chomp.split('').map(&:to_i) }
    nums = nums.transpose

    gamma, epsilon = [], []
    nums.each do |row|
      if row.count {|c| c == 1 } > row.size/2
        gamma << 1
        epsilon << 0
      else
        gamma << 0
        epsilon << 1
      end
    end
    epsilon = epsilon.join('').to_i(2)
    gamma = gamma.join('').to_i(2)

    puts epsilon * gamma
  end
#PART 2
# def call
#   @orig_nums = File.readlines('data.txt')
#   nums = @orig_nums.map {|i| i.to_s.chomp.split('') }
#   @nums = nums.transpose
#   indexes = (0..@orig_nums.size-1).to_a

#   oxygen = process_row(indexes, 0, 'oxygen')
#   co2 = process_row(indexes, 0, 'co2')
#   puts oxygen.to_i(2) * co2.to_i(2)

# end

# def one(arr, indexes)
#   count = indexes.count { |i| arr[i] == "1" }
#   (count >= indexes.size - count) ? "1" : "0"
# end

# def zero(arr, indexes)
#   count = indexes.count { |i| arr[i] == "0" }
#   (count <= indexes.size - count) ? "0" : "1"
# end

# def process_row(indexes, iter, val)
#   return if iter > 40
#   res = if val == 'oxygen'
#     one(@nums[iter], indexes)
#   else
#     zero(@nums[iter], indexes)
#   end

#   result = indexes.select { |a| @nums[iter][a] == res }
#   if result.size > 1
#     process_row(result, iter + 1, val)
#   else
#     return @orig_nums[result[0]]
#   end
# end

# # 00100
#     # 11110

#         # 10110
#         # 10111
#         # 10101

# # 01111
# # 00111
#     # 11100
#         # 10000
#     # 11001
# # 00010
# # 01010
end
AoC.new.call
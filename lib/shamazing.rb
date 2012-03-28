class Shamazing
  VERSION = '0.0.4'

  # Finds the longest continuous String in a SHA.
  #
  # sha - The String SHA to analyze.
  #
  # Returns a String.
  def self.string(sha)
    sha.
      downcase.
      scan(/[a-f]+/).
      sort{|a,b| b.length <=> a.length}.
      first
  end

  # Finds the longest continuous String in an Array of SHAs.
  #
  # shas - The Array of SHAs to analyze.
  # full - A Boolean: should we return the full SHA instead of the snippet?
  #
  # Returns the longest String.
  def self.string_from_array(shas,full=false)
    longest = shas.
      collect{|sha| string(sha)}.
      sort{|a,b| b.length <=> a.length}.
      first
    full ? shas.find{|sha| sha.match(/#{longest}/)} : longest
  end

  # Finds the longest continuous Integer in a SHA.
  #
  # sha - The String SHA to analyze.
  #
  # Returns an Integer.
  def self.integer(sha)
    sha.
      scan(/[0-9]+/).
      sort{|a,b| b.length <=> a.length}.
      first.to_i
  end

  # Finds the longest continuous Integer in an Array of SHAs.
  #
  # shas - The Array of SHAs to analyze.
  # full - A Boolean: should we return the full SHA instead of the snippet?
  #
  # Returns the longest Integer.
  def self.integer_from_array(shas,full=false)
    longest = shas.
      collect{|sha| integer(sha).to_s}.
      sort{|a,b| b.length <=> a.length}.
      first.to_i
    full ? shas.find{|sha| sha.match(/#{longest}/)} : longest
  end
end
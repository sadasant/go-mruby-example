class DirTests < MTest::Unit::TestCase

  def test_dir
    assert_equal(Class, Dir.class)
  end

  def test_entries
    entries = Dir.entries(".")
    assert_equal(Array, entries.class)
    assert(entries.include?("dir.rb"))
  end

  def test_pwd
    cwd = Dir.pwd.split("/")[-1]
    assert_equal(cwd, "mtests")
  end

  def test_chdir
    Dir.chdir("..")
    cwd = Dir.pwd.split("/")[-1]
    assert_equal(cwd, "go-mruby-example")
    Dir.chdir("mtests")
    cwd = Dir.pwd.split("/")[-1]
    assert_equal(cwd, "mtests")
  end

end

raise if MTest::Unit.new.run != 0

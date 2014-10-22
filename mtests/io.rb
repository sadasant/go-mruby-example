class IOTests < MTest::Unit::TestCase

  def test_io
    assert_equal(Class, IO.class)
  end

  def test_file_exists
    filename = "io.rb"
    assert(File.exists?(filename))
  end

  def test_binread
    filename = "../build_config.rb"
    contents =
'MRuby::Build.new do |conf|
  toolchain :gcc

  enable_debug

  conf.gem \'mrbgems/mruby-dir\'
  conf.gem \'mrbgems/mruby-io\'
  conf.gem \'mrbgems/mruby-mtest\'

  conf.gembox \'default\'
end
'

    assert(File.exists?(filename))

    file = File.open(filename, "rb")
    assert_equal(contents, file.read)
  end

end

raise if MTest::Unit.new.run != 0

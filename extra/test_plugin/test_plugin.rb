Plugin.create(:test) do
  onperiod do
    Plugin.call(:update, nil, [Mikutter::System::Message.new(description: "test description")])
  end
end


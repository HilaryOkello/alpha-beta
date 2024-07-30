class CreateHealthfacilities < ActiveRecord::Migration[7.1]
  def change
    create_table :healthfacilities do |t|

      t.timestamps
    end
  end
end

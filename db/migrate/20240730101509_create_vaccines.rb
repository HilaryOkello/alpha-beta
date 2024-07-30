class CreateVaccines < ActiveRecord::Migration[7.1]
  def change
    create_table :vaccines do |t|

      t.string :username
      t.string :password_digest
      t.integer :vaccine_id
      t.integer :manufacturer_id
      t.integer :distributor_id
      t.integer :healthfacility_id

      t.timestamps
    end
  end
end

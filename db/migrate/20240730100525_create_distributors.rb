class CreateDistributors < ActiveRecord::Migration[7.1]
  def change
    create_table :distributors do |t|

      t.timestamps
    end
  end
end

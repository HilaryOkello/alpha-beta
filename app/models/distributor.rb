class Distributor < ApplicationRecord
  belongs_to :vaccine
  belongs_to :manufacturer
  belongs_to :healthfacility

  has_secure_password
end

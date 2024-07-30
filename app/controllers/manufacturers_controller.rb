class ManufacturersController < ApplicationController
    def create
        manufacturer = Manufacturer.create(manufacturer_params)
        if manufacturer.valid?
            render json: manufacturer, status: :created
        else
            render json: { errors: manufacturer.errors.full_messages }, status: :unprocessable_entity
        end
    end

    def show
        manufacturer = Manufacturer.find_by(id: session[:manufacturer_id])
        if manufacturer
          render json: manufacturer
        else
          render json: { error: 'manufacturer not found' }, status: :unauthorized
        end
    end

    private

      def manufacturer_params
        params.permit(:username, :password, :password_confirmation)
      end
end

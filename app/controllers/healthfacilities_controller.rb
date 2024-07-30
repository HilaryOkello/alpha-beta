class HealthfacilitiesController < ApplicationController
    def create
        healthfacilities = Healthfacilitie.create(healthfacilities_params)
        if healthfacilities.valid?
            render json: healthfacilities, status: :created
        else
            render json: { errors: healthfacilities.errors.full_messages }, status: :unprocessable_entity
        end
    end

    def show
        healthfacilitie = Healthfacilitie.find_by(id: session[:healthfacilitie_id])
        if healthfacilities
          render json: healthfacilities
        else
          render json: { error: 'healthfacilities not found' }, status: :unauthorized
        end
    end

    private

      def healthfacilitie_params
        params.permit(:username, :password, :password_confirmation)
      end
end

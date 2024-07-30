class DistributorsController < ApplicationController
    def create
        distrubutor = Distrubutor.create(distributor_params)
        if distrubutor.valid?
            render json: distributor, status: :created
        else
            render json: { errors: distributor.errors.full_messages }, status: :unprocessable_entity
        end
    end

    def show
        distributor = Distributor.find_by(id: session[:distributor_id])
        if distributor
          render json: distributor
        else
          render json: { error: 'distributor not found' }, status: :unauthorized
        end
    end

    private

      def distributor_params
        params.permit(:username, :password, :password_confirmation)
      end
end

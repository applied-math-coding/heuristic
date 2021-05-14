module TargetPlot
import Base.MathConstants as cons
import Plots as plt
plt.pyplot()

function f(x, y)
    return -20.0 * exp(-0.2 * sqrt(0.5 * (x^2 + y^2))) -
  exp(0.5 * (cos(2.0 * pi * x) + cos(2.0 * pi * y))) +  cons.e + 20.0
  #  ((1.0 - y) * sin(x))^2 + (y * (2.0 - x))^2
end

const plot = function ()
    x = -5.0:0.1:5.0
    y = -5.0:0.1:5.0
    # plt.heatmap(x, y, f)
    plt.plot(x, y, f, st=:wireframe)
    plt.png("my-plots/wireframe.png")
end

end